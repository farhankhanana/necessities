package rabbitmq

import (
	"fmt"
	"strings"
	"time"

	"github.com/gat/necessities/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	globalAMQPURI           string
	globalConnection        *amqp.Connection
	globalChannel           *amqp.Channel
	globalTimeout           time.Duration
	globalProjectName       string
	globalProjectModule     string
	queueRPCServerName      string
	defaultQueueNameFormat  = "%s_%s_rpc"
	intervalDelay           = 5
	gracefulShutdownChannel = make(chan struct{}, 1)
	stopConsumingChannel    = make(chan struct{}, 1)
)

const (
	defaultTimeout = 40 * time.Second
)

type Exchange struct {
	ExchangeKey string
	RoutingKey  string
	QueueName   string
}

// GenerateQueueName generates queue name with custom format if provided.
// Format must have count total of string placeholders (%s) equal with len of `formatValue`.
// If there are no string placeholder, or string placeholders doesn't have equal value with len `formatValue`, default placeholder will be applied.
//
// default queue name is `%s_%s_rpc` if not overwritten.
func GenerateQueueName(format string, formatValue ...string) string {
	if len(format) < 1 || strings.Count(format, "%s") != len(formatValue) {
		format = defaultQueueNameFormat
	}
	for _, s := range formatValue {
		format = strings.Replace(format, "%s", s, 1)
	}
	return format
}

// OverwriteDefaultQueueNameFormat overwrites default format for queue name. New format must have at least 2 string placeholders (%s).
// Returns true if new format has 2 string placeholders, otherwise returns false and default value will be not overwrite.
func OverwriteDefaultQueueNameFormat(format string) bool {
	if strings.Count(format, "%s") >= 2 {
		defaultQueueNameFormat = format
		return true
	}
	return false
}

// InitRabbitMQ initiates Rabbit MQ service.
// The function needs project name, project module, timeout for response and also
// amqpURI consist of {protocol://user:password@host:port/parameter} of running Rabbit MQ.
// Run this first in your init config or main.go.
// Returns err that occurs while initializing RabbitMQ.
func InitRabbitMQ(projectName, projectModule, amqpURI string, timeout time.Duration) (err error) {
	logger := logger.NewLogger("")

	err = openGlobalConnection(amqpURI)
	if err != nil {
		logger.LogError("failed to connect to RabbitMQ", err.Error())
		return err
	}

	err = openGlobalChannel()
	if err != nil {
		defer closeActiveConnections()
		logger.LogError("failed to create channel", err.Error())
		return err
	}

	queueRPCServerName = GenerateQueueName("", projectName, projectModule)
	_, err = globalChannel.QueueDeclare(
		queueRPCServerName, // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		logger.LogError("failed to declare a queue consumer", err.Error())
		defer closeActiveConnections()
		return err
	}

	if timeout == 0 {
		globalTimeout = defaultTimeout
	} else {
		globalTimeout = timeout
	}

	globalAMQPURI = amqpURI
	globalProjectName = projectName
	globalProjectModule = projectModule
	logger.LogInfo("finish init rabbitMQ with queue name:", queueRPCServerName)

	// observe connection
	observeGlobalConnection()
	return nil
}

func openGlobalConnection(amqpURI string) (err error) {
	globalConnection, err = amqp.Dial(amqpURI)
	if err != nil {
		return err
	}
	return nil
}

func openGlobalChannel() (err error) {
	if globalConnection == nil || globalConnection.IsClosed() {
		err = openGlobalConnection(globalAMQPURI)
		if err != nil {
			return err
		}
	}

	if globalChannel == nil || globalChannel.IsClosed() {
		globalChannel, err = globalConnection.Channel()
		if err != nil {
			return err
		}
	}
	return nil
}

// closeActiveConnections closes active connection and channel.
func closeActiveConnections() (err error) {
	logger := logger.NewLogger("")
	if !globalChannel.IsClosed() {
		if err = globalChannel.Close(); err != nil {
			logger.LogError("close channel", err.Error())
			return err
		}
	}

	if globalConnection != nil && !globalConnection.IsClosed() {
		if err = globalConnection.Close(); err != nil {
			logger.LogError("close connection", err.Error())
			return err
		}
	}

	return err
}

// GracefulShutdown shutdowns global channel and global connection gracefully.
func GracefulShutdown() error {
	logger := logger.NewLogger("")
	gracefulShutdownChannel <- struct{}{}
	stopConsumingChannel <- struct{}{}

	logger.LogInfo("Waiting for all messages to be processed...")
	messageWG.Wait()
	logger.LogInfo("All messages have been processed. Proceeding to close connection...")

	return closeActiveConnections()
}

// observeGlobalConnection observes rabbit mq connection and channel. When connection or channel is close this function
// will be triggered and retry to reconnect.
func observeGlobalConnection() {
	logger := logger.NewLogger("")
	intervalDelayDuration := time.Duration(intervalDelay) * time.Second
	go func() {
		select {
		case <-gracefulShutdownChannel:
			logger.LogInfo("rabbitMQ is shutting down gracefully")
			return
		case err := <-globalConnection.NotifyClose(make(chan *amqp.Error)):
			logger.LogWarn("rabbitMQ global connection is closed unexpectedly", err.Error())
		case err := <-globalChannel.NotifyClose(make(chan *amqp.Error)):
			logger.LogWarn("rabbitMQ global channel is closed unexpectedly", err)
		}

		logger.LogInfo("closing remaining connection and channel...")
		_ = closeActiveConnections()

		logger.LogInfo(fmt.Sprintf("rabbitMQ is retrying to reconnect in %ds...", intervalDelay))
		time.Sleep(intervalDelayDuration)

		for err := InitRabbitMQ(globalProjectName, globalProjectModule, globalAMQPURI, globalTimeout); err != nil; err = InitRabbitMQ(globalProjectName, globalProjectModule, globalAMQPURI, globalTimeout) {
			logger.LogError("rabbitMQ error occurs while retrying to reconnect", err.Error())
			logger.LogInfo(fmt.Sprintf("rabbitMQ is retrying to reconnect in %ds...", intervalDelay))
			time.Sleep(intervalDelayDuration)
		}
		logger.LogInfo("rabbitMQ is reconnected")

		logger.LogInfo(fmt.Sprintf("rabbitMQ is re-initiating ConsumeRPC function in %ds...", intervalDelay))
		time.Sleep(intervalDelayDuration)

		go ConsumeRPC(globalRPCHandler)

		logger.LogInfo("re-initiating consumerRPC is done")
	}()
}

package rabbitmq

import (
	"encoding/json"
	"fmt"
	"sync"

	lgr "github.com/gat/necessities/logger"
	"github.com/gat/necessities/model"
	rsp "github.com/gat/necessities/response"
	"github.com/gat/necessities/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	globalRPCHandler func(delivery amqp.Delivery) rsp.Response
	messageWG        sync.WaitGroup
)

// ConsumeRPC serves as server consumer/subscriber of topic rpc needed from this module.
// Since v2.0.11 this function didn't call goroutine internally.
// So, you need to run this function in goroutine if you want to consume message from RabbitMQ in background.
func ConsumeRPC(rpcHandler func(delivery amqp.Delivery) rsp.Response) {
	logger := lgr.NewLogger("")
	messages, err := globalChannel.Consume(
		queueRPCServerName, // queue
		"",                 // consumer
		false,              // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	if err != nil {
		logger.LogFatal("failed to start RPC consumer", err.Error())
	}

	globalRPCHandler = rpcHandler
	consuming := true

	for consuming {
		select {
		case <-stopConsumingChannel:
			consuming = false
			break
		case m := <-messages:
			messageWG.Add(1)
			go func(message amqp.Delivery) {
				defer messageWG.Done()
				loggerRPC := lgr.NewLogger(message.CorrelationId)
				loggerRPC.LogInfo(
					fmt.Sprintf("received a RPC request with request_type : '%s'", message.Type),
					utils.AnyToMapStringInterface(message.Body),
				)
				var response rsp.Response

				// process rpc request if message body is not empty
				if len(message.Body) > 0 {
					response = rpcHandler(message)
				}

				exchangeInfo := Exchange{
					RoutingKey: message.ReplyTo,
				}

				responseByte, _ := json.Marshal(response)
				mq := &model.MessageQueue{
					CorrelationID: message.CorrelationId,
					ContentType:   message.ContentType,
					RequestType:   message.Type,
					Value:         responseByte,
				}

				err = PublishMessage(exchangeInfo, mq)
				if err != nil {
					loggerRPC.LogError(
						fmt.Sprintf("publish message failed : %s", err.Error()),
						utils.AnyToMapStringInterface(message.Body),
					)
				}
				// turn on manual acknowledge
				err = message.Ack(false)
				if err != nil {
					loggerRPC.LogError(
						fmt.Sprintf("acknowledge message failed : %s", err.Error()),
						utils.AnyToMapStringInterface(message.Body),
					)
				}
			}(m)
		}
	}
}

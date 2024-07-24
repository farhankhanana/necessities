package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gat/necessities/logger"
	"github.com/gat/necessities/model"
	rsp "github.com/gat/necessities/response"
	"github.com/gat/necessities/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	queueExpiredDuration = 120000 // in millisecond
)

// SendRPC sends or calls message request or function from other module using RPC.
// Accept MessageQueue Model from ces-utilities/model/v2
// Field explanations from MessageQueue for SendRPC are:
//   - `CorrelationID` optional, automatically UUID generated if is empty
//   - `RequestType` mandatory, worker name or case name from other service
//   - `ContentType` optional, automatically set to 'application/json' if it's empty
//   - `Header` conditional, set it if needed
//   - `Value` mandatory, as value to be set as argument/request to other service
//   - `Timeout` optional, using time.Duration. If it's empty then it will use `globalTimeout` from RabbitMQ initiation
//
// For field `Value` could be struct, []byte, even nil.
//
// `targetModule` works as name of other service/module.
func SendRPC(mq model.MessageQueue, targetModule string) rsp.Response {
	var err error
	var tempChannel *amqp.Channel
	var q amqp.Queue
	var messages <-chan amqp.Delivery

	if len(mq.CorrelationID) <= 0 {
		mq.CorrelationID = utils.UUIDGenerator()
	}
	if mq.Timeout <= 0 {
		mq.Timeout = globalTimeout
	}

	logger := logger.NewLogger(mq.CorrelationID)
	response := rsp.Response{
		StatusCode: 500,
	}

	tempChannel, err = globalConnection.Channel()
	if err != nil {
		logger.LogError("failed to open temp channel", err)
		response.ResponseData = rsp.GenerateJsonErrorResponse(err.Error(), nil, nil, rsp.RCInternalServerError, "")
		return response
	}

	queueSendRPCName := fmt.Sprintf(
		"%s_%s_%s_%s",
		globalProjectName,
		targetModule,
		mq.RequestType,
		mq.CorrelationID,
	)
	// declare temporary queue
	q, err = tempChannel.QueueDeclare(
		queueSendRPCName, // name
		true,             // durable
		true,             // delete when unused
		true,             // exclusive
		false,            // noWait
		amqp.Table{
			"x-expires": queueExpiredDuration,
		}, // arguments
	)
	if err != nil {
		logger.LogError("failed to declare a temp queue consumer", err)
		response.ResponseData = rsp.GenerateJsonErrorResponse(err.Error(), nil, nil, rsp.RCInternalServerError, "")
		return response
	}

	// create function to clean up queue and channel
	cleanUp := func() {
		if tempChannel == nil {
			return
		} else if tempChannel.IsClosed() {
			return
		}
		_, err = tempChannel.QueueDelete(
			q.Name,
			false, // ifUnused
			false, // ifEmpty
			true,  // noWait
		)
		if err != nil {
			logger.LogError("failed to delete queue", err.Error())
		}
		if err = tempChannel.Close(); err != nil {
			logger.LogError("failed to close channel", err.Error())
		}
	}

	defer func() {
		go cleanUp()
	}()

	// declare temporary consumer
	messages, err = tempChannel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logger.LogError("failed to register a queue consumer", err)
		response.ResponseData = rsp.GenerateJsonErrorResponse(err.Error(), nil, nil, rsp.RCInternalServerError, "")
		return response
	}

	// set exchange information
	exchangeInfo := Exchange{
		RoutingKey: GenerateQueueName("", globalProjectName, targetModule),
		QueueName:  q.Name,
	}
	logger.LogInfo(fmt.Sprintf("publish message using RPC to rabbit mq with routing key %s", exchangeInfo.RoutingKey),
		utils.AnyToMapStringInterface(mq),
	)

	err = PublishMessage(exchangeInfo, &mq)
	if err != nil {
		logger.LogError("publish message", err)
		response.ResponseData = rsp.GenerateJsonErrorResponse(err.Error(), nil, nil, rsp.RCInternalServerError, "")
		return response
	}

	//for message := range messages {
	//	if mq.CorrelationID == message.CorrelationId {
	//		json.Unmarshal(message.Body, &response)
	//		break
	//	}
	//}
mqLoop:
	for {
		logger.LogInfo(fmt.Sprintf("loop-select waiting for response from %s...", exchangeInfo.RoutingKey))
		select {
		case message := <-messages:
			if mq.CorrelationID == message.CorrelationId {
				logger.LogInfo(fmt.Sprintf("got message reply from %s", exchangeInfo.RoutingKey),
					utils.AnyToMapStringInterface(message.Body),
				)
				err = json.Unmarshal(message.Body, &response)
				if err != nil {
					logger.LogError("error unmarshal message body", err)
					response.ResponseData = rsp.GenerateJsonErrorResponse(err.Error(), nil, nil, rsp.RCInternalServerError, "")
					return response
				}
				break mqLoop
			}
		case <-time.After(mq.Timeout):
			logger.LogWarn(
				fmt.Sprintf("waiting message reply timeout from %s", exchangeInfo.RoutingKey),
			)
			response.StatusCode = rsp.HTTPRequestTimeout
			response.ResponseData = rsp.GenerateJsonErrorResponse("rabbit mq timeout", nil, nil, rsp.RCConnectionTimeout, "")
			break mqLoop
		}
	}

	return response
}

// PublishMessage publish message to queue based on exchange information (e.g. queue name, exchange key, routing key).
//
// If this func intended for RPC between services, then `Exchange.queue_name` must be filled with queue name of publisher (as reply-to),
// and `Exchange.routing_key` filled with queue name of target module or service.
//
// If this func intended for send to downstream, then `Exchange.queue_name` must be empty. And
// `Exchange.routing_key` filled with queue name of topic stream.
func PublishMessage(exchangeInfo Exchange, mq *model.MessageQueue) error {
	if len(mq.CorrelationID) <= 0 {
		mq.CorrelationID = utils.UUIDGenerator()
	}
	logger := logger.NewLogger(mq.CorrelationID)

	if len(mq.ContentType) <= 0 {
		mq.ContentType = "application/json"
	}

	timestampNow, _ := utils.TimestampNow(false, "")

	publishData := amqp.Publishing{
		CorrelationId: mq.CorrelationID,
		ContentType:   mq.ContentType,
		Timestamp:     timestampNow,
		Type:          mq.RequestType,
	}

	if len(exchangeInfo.QueueName) > 0 {
		publishData.ReplyTo = exchangeInfo.QueueName
	}

	// switch case for type assertion header
	switch v := mq.Headers.(type) {
	case map[string]interface{}:
		publishData.Headers = v
	case nil:
		publishData.Headers = nil
	default:
		var newMap map[string]interface{}
		vByte, _ := json.Marshal(v)
		err := json.Unmarshal(vByte, &newMap)
		if err != nil {
			logger.LogError("error unmarshal vByte", err)
			return err
		}
		publishData.Headers = newMap
	}

	// switch case for type assertion body
	switch v := mq.Value.(type) {
	case []byte:
		publishData.Body = v
	case nil:
		publishData.Body = nil
	default:
		vByte, _ := json.Marshal(v)
		publishData.Body = vByte
	}

	ctx, cancel := context.WithTimeout(context.Background(), mq.Timeout)
	defer cancel()

	newChannel, err := globalConnection.Channel()
	if err != nil {
		return err
	}
	defer newChannel.Close()

	err = newChannel.PublishWithContext(
		ctx,
		"",                      // exchange
		exchangeInfo.RoutingKey, // routing key
		false,                   // mandatory
		false,                   // immediate
		publishData,
	)
	return err
}

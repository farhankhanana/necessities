package model

import (
	"sync"
	"time"
)

// MessageQueue defines single object as request or response of message queue.
// Can be used for other message queue, such as Kafka, RabbitMQ, KubeMQ, and others
type MessageQueue struct {
	Headers       interface{}   `json:"headers,omitempty"`
	Value         interface{}   `json:"value,omitempty"`
	RequestID     string        `json:"request_id,omitempty"`
	CorrelationID string        `json:"correlation_id,omitempty"`
	RequestType   string        `json:"request_type,omitempty"`
	ContentType   string        `json:"content_type,omitempty"`
	Key           string        `json:"key,omitempty"`
	NeedResponse  bool          `json:"need_response,omitempty"`
	Timeout       time.Duration `json:"timeout,omitempty"`
}

// Queue defines single object in message queue table.
type Queue struct {
	IsDone   chan bool
	Request  MessageQueue
	Response MessageQueue
}

// MessageQueueTable defines queue table request and response for
// message queue.
type MessageQueueTable struct {
	QueueTable map[string]*Queue
	Mutex      sync.Mutex
}

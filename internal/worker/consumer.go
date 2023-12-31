package worker

import (
	"fmt"
	"github.com/adjust/rmq/v5"
	"log"
	"time"
	"worker/internal/worker/handlers"
)

type Consumer struct {
	name      string
	queueName string
	count     int
	before    time.Time
}

func NewConsumer(tag int, queueName string) *Consumer {
	return &Consumer{
		name:      fmt.Sprintf("consumer_%d", tag),
		queueName: queueName,
		count:     0,
		before:    time.Now(),
	}
}

func (consumer *Consumer) Consume(delivery rmq.Delivery) {

	payload := delivery.Payload()
	switch consumer.queueName {
	case queueThings:
		handlers.HandleThings(consumer.queueName, payload)
	}

	//log.Printf("start consumer for queue: %v with payload: %s", consumer.queueName, payload)
	time.Sleep(consumeDuration)

	consumer.count++
	if consumer.count%reportBatchSize == 0 {
		duration := time.Now().Sub(consumer.before)
		consumer.before = time.Now()
		perSecond := time.Second / (duration / reportBatchSize)
		log.Printf("%s consumed %d %s %d", consumer.name, consumer.count, payload, perSecond)
	}

	if consumer.count%reportBatchSize > 0 {
		if err := delivery.Ack(); err != nil {
			debugf("failed to ack %s: %s", payload, err)
		} else {
			debugf("acked %s", payload)
		}
	} else { // reject one per batch
		if err := delivery.Reject(); err != nil {
			debugf("failed to reject %s: %s", payload, err)
		} else {
			debugf("rejected %s", payload)
		}
	}
}

package worker

import (
	"github.com/adjust/rmq/v5"
	"log"
	"time"
)

type BatchConsumer struct {
	tag string
}

func NewBatchConsumer(tag string) *BatchConsumer {
	return &BatchConsumer{tag: tag}
}

func (consumer *BatchConsumer) Consume(batch rmq.Deliveries) {
	payloads := batch.Payloads()
	debugf("start consume %q", payloads)
	time.Sleep(consumeDuration)

	log.Printf("%s consumed %d: %s", consumer.tag, len(batch), batch[0])
	errors := batch.Ack()
	if len(errors) == 0 {
		debugf("acked %q", payloads)
		return
	}

	for i, err := range errors {
		debugf("failed to ack %q: %q", batch[i].Payload(), err)
	}
}

package worker

import (
	"fmt"
	"log"
	"time"
)

func PublishToQueues() {

	things, err := GetOpenQueue(queue1)
	if err != nil {
		panic(err)
	}
	foobars, err := GetOpenQueue(queue2)
	if err != nil {
		panic(err)
	}

	var before time.Time
	for i := 0; i < 10000000; i++ {
		delivery := fmt.Sprintf("delivery %d", i)
		if err := things.Queue.Publish(delivery); err != nil {
			log.Printf("failed to publish: %s", err)
		}

		if i%batchSize == 0 {
			duration := time.Now().Sub(before)
			before = time.Now()
			perSecond := time.Second / (duration / batchSize)
			log.Printf("produced %d %s %d", i, delivery, perSecond)
			if err := foobars.Queue.Publish("Testing Batching"); err != nil {
				log.Printf("failed to publish: %s", err)
			}
		}
	}
}

package worker

import (
	"fmt"
	"log"
	"time"
)

func PublishToQueues() {

	things, err := GetOpenQueue(queueThings)
	if err != nil {
		panic(err)
	}
	//things.Purge()

	foobars, err := GetOpenQueue(queueFoobars)
	if err != nil {
		panic(err)
	}
	//foobars.Purge()

	var before time.Time
	for i := 0; i < 100; i++ {
		delivery := fmt.Sprintf("delivery %d", i)
		if err := things.Queue.Publish(delivery); err != nil {
			log.Printf("failed to publish: %s", err)
		}

		if i%batchSize == 0 {
			duration := time.Now().Sub(before)
			before = time.Now()
			perSecond := time.Second / (duration / batchSize)
			log.Printf("produced %d %s %d", i, delivery, perSecond)
			// Try Queue.PublishBytes() in case of []string
			if err := foobars.Queue.Publish(delivery); err != nil {
				log.Printf("failed to publish: %s", err)
			}
		}
	}
}

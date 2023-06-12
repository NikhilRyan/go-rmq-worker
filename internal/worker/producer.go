package worker

import (
	"log"
)

func Push(queueName, payload string) {

	queue, err := connection.OpenQueue(queueName)
	if err != nil {
		log.Printf("failed to open queue (%v): %s", queueName, err)
		return
	}

	if publishErr := queue.Publish(payload); publishErr != nil {
		log.Printf("failed to publish in queue (%v): %s", queueName, publishErr)
	}
}

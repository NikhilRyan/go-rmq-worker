package handlers

import "log"

func HandleThings(queueName, payload string) {
	log.Printf("Payload received for queue: %v with data: %s", queueName, payload)
}

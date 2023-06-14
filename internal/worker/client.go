package worker

import (
	"fmt"
	"github.com/adjust/rmq/v5"
	"log"
	"sync"
)

var (
	connection rmq.Connection
	once       sync.Once
)

func InitializeConnection() (err error) {
	errChan := make(chan error, 10)
	go logErrors(errChan)

	once.Do(func() {
		connection, err = rmq.OpenConnection("myBroker", "tcp", "localhost:6379", 2, errChan)
		if err != nil {
			log.Print("Unable to create redis connection for queues: ", err)
		}
	})

	return err
}

func GetConnection() rmq.Connection {
	return connection
}

func StartConsumers() (errors []error) {

	// Queues
	queueConfigs := GetAllQueueWithConfig()

	for _, config := range queueConfigs {
		if !config.IsActive {
			log.Printf("Inactive queue %v, not initialised.", config.Name)
			continue
		}

		queue, OpenQueueErr := connection.OpenQueue(config.Name)
		if OpenQueueErr != nil {
			log.Printf("Unable to open queue (%v), error: %v", config.Name, OpenQueueErr)
			errors = append(errors, OpenQueueErr)
			continue
		}

		if startErr := queue.StartConsuming(config.PrefetchLimit, config.PollDuration); startErr != nil {
			log.Printf("Unable to start consumer for queue (%v), error: %v", config.Name, startErr)
			errors = append(errors, startErr)
			continue
		}

		// TODO: Make multiple consumers in batch also
		if config.BatchingEnabled {
			if _, abcErr := queue.AddBatchConsumer(config.Name, config.BatchSize, config.BatchTimeout, NewBatchConsumer(config.Name)); abcErr != nil {
				log.Printf("Unable to add batch consumer for queue (%v), error: %v", config.Name, abcErr)
				errors = append(errors, abcErr)
				continue
			}
		} else {
			for i := 0; i < config.NumWorkers; i++ {
				consumerName := fmt.Sprintf(consumerNameString, config.Name, i)
				if _, addConsumerErr := queue.AddConsumer(consumerName, NewConsumer(i, config.Name)); addConsumerErr != nil {
					log.Printf("Unable to add consumer for queue (%v), error: %v", config.Name, addConsumerErr)
					errors = append(errors, addConsumerErr)
					break
				}

			}
		}
	}

	return
}

func logErrors(errChan <-chan error) {
	for err := range errChan {
		switch err := err.(type) {
		case *rmq.HeartbeatError:
			if err.Count == rmq.HeartbeatErrorLimit {
				log.Print("heartbeat error (limit): ", err)
			} else {
				log.Print("heartbeat error: ", err)
			}
		case *rmq.ConsumeError:
			log.Print("consume error: ", err)
		case *rmq.DeliveryError:
			log.Print("delivery error: ", err.Delivery, err)
		default:
			log.Print("other error: ", err)
		}
	}
}

func debugf(format string, args ...interface{}) {
	if shouldLog {
		log.Printf(format, args...)
	}
}

package worker

import "time"

// Queue Names
const (
	queueThings   = "things"
	queueFoobars  = "foobars"
	queueDoRandom = "dorandom"
)

// Consumer
const (
	prefetchLimit = 1000
	pollDuration  = 100 * time.Millisecond
	numConsumers  = 5
	batchSize     = 10
	batchTimeout  = time.Second

	reportBatchSize = 10000
	consumeDuration = time.Millisecond
	shouldLog       = false

	consumerNameString = "queue_%v_consumer_%d"
)

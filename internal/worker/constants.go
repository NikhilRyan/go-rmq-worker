package worker

import "time"

// Queue Names
const (
	queue1 = "queue1"
	queue2 = "queue2"
	queue3 = "queue3"
)

// Consumer
const (
	prefetchLimit = 1000
	pollDuration  = 100 * time.Millisecond
	numConsumers  = 5
	batchSize     = 111
	batchTimeout  = time.Second

	reportBatchSize = 10000
	consumeDuration = time.Millisecond
	shouldLog       = false
)

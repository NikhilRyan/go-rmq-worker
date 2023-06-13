package worker

import "time"

// Queue Names
const (
	queue1 = "things"
	queue2 = "foobars"
	queue3 = "dorandom"
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

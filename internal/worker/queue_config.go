package worker

import "time"

type QueueConfig struct {
	Name          string
	NumWorkers    int
	PrefetchLimit int64
	PollDuration  time.Duration
	IsActive      bool

	BatchingEnabled bool
	BatchSize       int64
	BatchTimeout    time.Duration
}

func GetNewQueueConfig(name string, numWorkers int, prefetchLimit int64, pollDuration time.Duration, isActive bool) QueueConfig {
	return QueueConfig{
		Name:          name,
		NumWorkers:    numWorkers,
		PrefetchLimit: prefetchLimit,
		PollDuration:  pollDuration,
		IsActive:      isActive,
	}
}

func (qc *QueueConfig) AddBatchConfig(enabled bool, size int64, timeout time.Duration) {
	qc.BatchingEnabled = enabled
	qc.BatchSize = size
	qc.BatchTimeout = timeout
}

func GetAllQueueWithConfig() (queues []QueueConfig) {

	// First Queue
	firstQueue := GetNewQueueConfig(queueThings, numConsumers, prefetchLimit, pollDuration, true)
	queues = append(queues, firstQueue)

	// Second Queue
	SecondQueue := GetNewQueueConfig(queueFoobars, numConsumers, prefetchLimit, pollDuration, true)
	SecondQueue.AddBatchConfig(true, batchSize, batchTimeout)
	queues = append(queues, SecondQueue)

	// Third Queue
	ThirdQueue := GetNewQueueConfig(queueDoRandom, prefetchLimit, 1, pollDuration, false)
	queues = append(queues, ThirdQueue)

	return
}

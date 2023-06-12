package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"worker/internal/worker"
)

func main() {

	initializeRedisQueueWorkers()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	defer signal.Stop(signals)

	<-signals // wait for signal
	go func() {
		<-signals // hard exit on second signal (in case shutdown gets stuck)
		os.Exit(1)
	}()

	<-worker.GetConnection().StopAllConsuming() // wait for all Consume() calls to finish
}

func initializeRedisQueueWorkers() {

	initError := worker.InitializeConnection()
	if initError != nil {
		log.Panicf("Unable to initialize redis queue, error: %v", initError)
	}

	consumerErrors := worker.StartConsumers()
	if len(consumerErrors) > 0 {
		log.Printf("Unable to initialize redis queue workers, errors: %v", consumerErrors)
	}
}

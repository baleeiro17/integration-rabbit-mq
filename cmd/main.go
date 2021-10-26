package main

import (
	"integration-rabbit-mq/internal/consumer"
	"integration-rabbit-mq/internal/publisher"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}

	go publisher.Send()

	wg.Add(1)

	go consumer.Receive()

	wg.Add(1)

	wg.Wait()

}

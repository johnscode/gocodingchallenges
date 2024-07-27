package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var buffer = make(chan int, 5)

func produce(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		buffer <- i
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
	fmt.Println("producer done")
}

func consume(wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range buffer {
		fmt.Println(data)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
	}
	fmt.Println("consumer done")
}

func main() {
	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup
	producerWg.Add(1)
	go produce(&producerWg)
	go func() {
		producerWg.Wait()
		close(buffer)
		fmt.Println("closed channel")
	}()
	consumerWg.Add(1)
	go consume(&consumerWg)
	consumerWg.Wait()
	fmt.Println("done")
}

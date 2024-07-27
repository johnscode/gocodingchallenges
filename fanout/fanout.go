package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// produce is simulating our single input as a channel
func produce(id int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- rand.Intn(20)
		}
		fmt.Printf("producer %d done\n", id)
		close(ch) // this is important!!!
	}()
	return ch
}

func worker(id int, jobs chan int, wg *sync.WaitGroup) {
	for value := range jobs {
		odd := "even"
		if (value & 1) == 1 {
			odd = "odd"
		}
		fmt.Printf("worker: %d, got %d is %s\n", id, value, odd)
	}
	wg.Done()
}

func main() {
	inputCh := produce(1)

	numWorkers := 3
	jobs := make(chan int)

	// split input into individual jobs
	go func() {
		for value := range inputCh {
			jobs <- value
		}
		close(jobs) // this is important!!!
	}()

	// fan-out
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}
	wg.Wait()

	fmt.Println("done")
}

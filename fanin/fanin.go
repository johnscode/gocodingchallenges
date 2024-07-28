package main

import (
	"fmt"
	"sync"
)

// produce is used to simulate the different data sources
func produce(id int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- id*10 + i
		}
		fmt.Printf("producer %d done\n", id)
		close(ch) // this is important!!!
	}()
	return ch
}

func fanin(inputs ...chan int) chan int {
	output := make(chan int)

	var wg sync.WaitGroup
	for i, input := range inputs {
		wg.Add(1)
		go func() {
			for value := range input {
				output <- value
			}
			fmt.Printf("done merging source %d\n", i)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(output) // this is important!!!
	}()

	return output
}

func main() {
	input1 := produce(0)
	input2 := produce(1)

	result := fanin(input1, input2)

	done := make(chan bool)
	go func() {
		for value := range result {
			fmt.Printf("got %d\n", value)
		}
		close(done)
	}()
	<-done
	fmt.Println("done")
}

package main

import (
	"fmt"
	"math/rand"
)

func produce(num int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < num; i++ {
			out <- rand.Intn(100)
		}
		close(out)
	}()
	return out
}

func double(input <-chan int) chan int {
	out := make(chan int)
	go func() {
		for value := range input {
			out <- value * 2
		}
		close(out)
	}()
	return out
}

func filterBelow10(input <-chan int) chan int {
	out := make(chan int)
	go func() {
		for value := range input {
			if value > 10 {
				out <- value
			}
		}
		close(out)
	}()
	return out
}

func print(input <-chan int) {
	for value := range input {
		fmt.Printf("value is %d\n", value)
	}
}

func main() {

	print(filterBelow10(double(produce(10))))

}

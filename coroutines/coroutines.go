package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Println("Goroutine", i, "started")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)               // Increment the counter
		go printNumbers(i, &wg) // Start a new goroutine
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines finished")

	// time.Sleep(4 * time.Second)
}

package main

import (
	"fmt"
	"time"
)

func processData(messageChan chan int) {
	// Simulate data processing
	for num := range messageChan {
		fmt.Println("Processing data:", num)
		time.Sleep(200 * time.Millisecond)
	} // Simulate a delay
}

func sum(result chan int, a, b int) {
	// Calculate the sum
	sum := a + b
	result <- sum // Send the result to the channel
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("p")

}

/*
	func main() {
		// Create a channel
		messageChan := make(chan int)
		go processData(messageChan)
		for {
			messageChan <- rand.Intn(100) // Send a random number to the channel
		}
		time.Sleep(2 * time.Second) // Simulate a delay

		// Start a goroutine to process data
	}

func main() {
	result := make(chan int)
	go sum(result, 5, 10)    // Start a goroutine to calculate the sum
	sum := <-result          // Receive the result from the channel   this is also blocking
	fmt.Println("Sum:", sum) // Print the result
}
*/

func main() {
	emailChan := make(chan string, 2)

	emailChan <- "drftgyh"
	emailChan <- "ftghj"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)
	close(emailChan)
}

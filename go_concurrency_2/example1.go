package main

import (
	"fmt"
	"time"
)

// Producer function
func producer(dataChannel chan int, done chan bool) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		dataChannel <- i        // Send data to the channel
		time.Sleep(time.Second) // Simulate time taken to produce
	}
	close(dataChannel) // Close the channel when done producing
	done <- true
}

// Consumer function
func consumer(dataChannel chan int, done chan bool) {
	for item := range dataChannel {
		fmt.Printf("Consuming: %d\n", item)
		time.Sleep(2 * time.Second) // Simulate time taken to consume
	}
	done <- true
}

func example1_main() {
	dataChannel := make(chan int)
	done := make(chan bool)

	// Start producer and consumer goroutines
	go producer(dataChannel, done)
	go consumer(dataChannel, done)

	// Wait for both goroutines to finish
	<-done
	<-done

	fmt.Println("Producer and Consumer have finished.")
}

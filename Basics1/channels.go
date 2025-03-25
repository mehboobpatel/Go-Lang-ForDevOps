package main

import (
	"fmt"
	"time"
)

func main() {
	u := make(chan int)  // unbuffered channel
	fmt.Println("u=", u) // memory address of u

	messages := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		messages <- "Hello, Channels!" //
	}()

	// Start a goroutine to send message asynchronously
	go func() {
		messages <- "Hello, inputs!" // No blocking since it's inside a goroutine
	}()

	fmt.Println("Waiting for message...")
	msg1 := <-messages // Receiving the first message
	fmt.Println("Received first message:", msg1)

	msg2 := <-messages // Receiving the second message
	fmt.Println("Received second message:", msg2)
}

package main

import (
	"fmt"
	"time"
)

// A Worker has task with instructions to be followed of sending string objects
func worker(ch chan string) {
	time.Sleep(2 * time.Second) // Simulating some work (2 sec delay)
	ch <- "Task1 Completed"     // Sending data into the channel (pipe)
	ch <- "Task2"
}

func worker2(ch chan string) {
	time.Sleep(2 * time.Second) // Simulating some work (2 sec delay)
	ch <- "Task3 Completed"
}

func main() {
	// Creating a channel 'ch' that allows sending and receiving string messages.
	ch := make(chan string) // Think of 'ch' as a pipeline for string data

	// Starting the worker function in a separate goroutine
	go worker(ch) // Calls the object sending worker to work through a channel

	go worker2(ch) // Calls the object sending worker to work through a channel

	// 'msg is variable' recieves data objects & implicitly becomes a string because 'ch' is a channel for string objects.
	//before this the channel belt will looklike (1stelemnt, 2ndlement.3rdelemtn)
	msg := <-ch
	//after this the channel belt will looklike (2ndlement.3rdelemtn)
	msg2 := <-ch
	//after this the channel belt will looklike (3rdelemtn)

	msg3 := <-ch
	//after this the channel belt will get empty ()

	fmt.Println(msg)
	fmt.Println(msg2)
	fmt.Println(msg3)
}

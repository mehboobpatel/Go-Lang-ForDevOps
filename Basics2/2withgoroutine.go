package main

import (
	"fmt"
	"time"
)

func printMessage(task string) {
	startTime := time.Now() // Capture start time
	fmt.Printf("%s started at %s\n", task, startTime.Format("15:04:05"))

	time.Sleep(1 * time.Second) // Simulating some work

	endTime := time.Now() // Capture end time
	fmt.Printf("%s ended at %s\n", task, endTime.Format("15:04:05"))
}

func main() {
	fmt.Println("Starting tasks...")

	// Launching each task in a separate goroutine
	go printMessage("Task 1")
	go printMessage("Task 2")
	go printMessage("Task 3")

	// Since every goroutine is running concurrently, the order of completion is not guaranteed. also the main function will not wait for the goroutines to complete.
	//Since the MAIN function is also a goroutine, it will not wait for the other goroutines to complete.
	//The main function will exit as soon as it completes its execution.
	//To prevent this, we can use a time.Sleep() function to wait for the goroutines to complete.

	// Wait to ensure all goroutines complete before exiting main Not a recommended way to wait for goroutines to complete
	time.Sleep(1 * time.Second) // if we increase time we might see the order of completion of tasks
	fmt.Println("All tasks completed.")
}

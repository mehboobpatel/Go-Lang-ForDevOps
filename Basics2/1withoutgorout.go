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
	fmt.Printf("%s ended at %s\n\n", task, endTime.Format("15:04:05"))
}

func main() {
	printMessage("Task 1")
	printMessage("Task 2")
	printMessage("Task 3")
}

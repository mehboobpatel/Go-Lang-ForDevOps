package main

import (
	"fmt"
	"time"
)

// Function that handles the processing of each order
func processOrder(i int, sendmssg chan string, sendid chan int) {
	time.Sleep(time.Second * 1)                       // Simulating processing time increase the time to see the magic
	sendmssg <- fmt.Sprintf("Packing Order %d", 10+i) //  Sending a message to textBelt via storing the returned value from Sprintf
	//fmt.Sprintf() → Returns a string that can be stored or used elsewhere.
	// fmt.printf() Prints to the console

	sendid <- i // Sending order ID to numberBelt

}

func main() {
	textBelt := make(chan string) // Conveyor belt for status messages
	numberBelt := make(chan int)  // Conveyor belt for order IDs

	// Simulating 5 orders being processed concurrently
	// Call the 'processOrder' function in a goroutine
	for i := 1; i <= 5; i++ {
		go processOrder(i, textBelt, numberBelt)
	}
	// After the above go function is processed both the belt are now filled with objects right to left (FIFO)
	//NB (5,1,3,4,2)         || TB (15,11,13,14,12)

	// Receiving processed orders
	for i := 1; i <= 5; i++ {
		//before recieivng the belt looks like (1,2,3,4,5,)
		message := <-textBelt   // Receiving message from the belt
		orderID := <-numberBelt // Receiving order ID

		//As soon as the data objects are recieved in variables(orderID, message)
		// the channels remove the value and moves forward just like in conveyor belt

		//So after first loop the belt will look like(2,3,4,5)

		// Print the status of each order
		fmt.Printf("Factory Status: %s | Processing Order ID: %d\n", message, orderID)
	}
}

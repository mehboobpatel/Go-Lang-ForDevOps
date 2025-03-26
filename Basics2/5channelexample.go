package main

import (
	"fmt"
	"time"
)

// Order represents a customer order with an ID and status
type Order struct {
	ID     int
	Status string
}

// processOrder simulates order processing and sends the status through a channel
func processOrder(orderID int, orderStatusChannel chan Order) {
	// Simulating some processing time (to mimic real-world delays like cooking food or assembling an order)
	time.Sleep(time.Second * 2)

	// Sending a new Order struct with the order ID and status "Completed" into the channel
	// `channel <- value` sends a value of type Order struct (i.e sending a whole struct into the channel).
	orderStatusChannel <- Order{ID: orderID, Status: "Completed"}
}

func main() {
	// Creating a channel to communicate order statuses between goroutines
	orderUpdateChannel := make(chan Order)

	// Launching multiple goroutines to process different orders concurrently
	for i := 1; i <= 5; i++ {
		// Each order is processed in a separate goroutine, allowing parallel execution
		go processOrder(i, orderUpdateChannel) // Passing the same channel to all goroutines
	}

	// Receiving and printing order statuses from the channel
	for i := 1; i <= 5; i++ {
		// Receiving processed order details from the channel
		// `value := <-channel` receives a value from the channel.
		processedOrder := <-orderUpdateChannel // Receiving the orders (i.e., Structs of Order type) from the channel
		//The := automatically infers the type based on what’s received from orderUpdateChannel.

		// Printing the received order status
		fmt.Printf("Order %d processed: %s\n", processedOrder.ID, processedOrder.Status)
	}
}

package main

import (
	"fmt"
	"time"
)

func outsidemain(outsidemain chan string) {
	time.Sleep(2 * time.Second)

	outsidemain <- "MESSAGE 3 FROM OUTSIDE MAIN!" // sending
}

func main() {
	//An Empty belt is created which will accept the objects of string type
	messagesbelt := make(chan string)

	test3 := "Message 1 WITHOUT PARAMS go inline function!"
	// Define a go inline/lambda function without parameters
	go func() {
		time.Sleep(2 * time.Second)

		messagesbelt <- test3 // sending
	}()

	// Define a go inline/lambda function with parameters
	go func(newname string) {
		time.Sleep(2 * time.Second)
		messagesbelt <- newname // sending
	}("Message2 WITH  PARAMS go inline function!")
	// Start a goroutine to send message asynchronously

	go outsidemain(messagesbelt)

	for i := 1; i <= 2; i++ {
		//Running the loop till two so that the first two objects are recieved
		//The belt will look like (onelement,second element,thirdelement) before the loop
		recvedmessage := <-messagesbelt
		//The belt will look like (onelement) after the loop

		fmt.Println("Received message:", i, recvedmessage)
	}

	fmt.Println("Waiting for message...")
	msg1 := <-messagesbelt // Receiving the first message
	//Recievening the remainging onelement from the belt
	//The belt will look like (empty) after this
	fmt.Println("Received first message:", msg1)

	// msg2 := <-messages // Receiving the second message
	// fmt.Println("Received second message:", msg2)

	// msg3 := <-messages // Receiving the second message
	// fmt.Println("Received second message:", msg3)
}

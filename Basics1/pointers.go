package main

import "fmt"

// Function that takes a pointer to an integer as a parameter
func updateValue(pointerToNumber *int) {
	fmt.Println("Value of pointerToNumber (Address stored in pointerToNumber):", pointerToNumber)
	fmt.Println("Address of pointerToNumber (Memory location of the pointer variable itself):", &pointerToNumber)

	// Dereferencing pointerToNumber and modifying the original variable
	*pointerToNumber = *pointerToNumber + 10
}

func main() {
	// Declaring an integer variable
	var number int = 20

	fmt.Println("Before updating, value of number:", number)
	fmt.Println("Memory address of number:", &number)

	// Creating a pointer variable that stores the address of 'number'
	var pointerToNumber *int = &number
	//pointerToNumber := &number     <---- SHORTHAND

	// Dereferencing pointerToNumber to get the actual value stored at the address
	fmt.Println("Dereferenced value of pointerToNumber:", *pointerToNumber)

	// Passing the address of 'number' to the function
	updateValue(&number)

	// Printing the updated value of 'number' after the function modifies it
	fmt.Println("After updating, value of number:", number)
}

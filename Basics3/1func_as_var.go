package main

import "fmt"

// Define a function 'y' that takes an integer 'y' as input has a return type as int
func y(y int) int {
	return y * y
}

func main() {
	// Print a message to indicate that the concept of "function as a variable" is being demonstrated
	fmt.Println("function as variable:")

	// Declare a variable 'v' that will hold a function. This function will take an integer as input and return an integer.
	// Here we are not defining the function body, we are just declaring the variable 'v' to hold a function.
	// variable v is not getting assigned to any function yet.
	var v func(int) string

	// Declare another variable 'z' to hold a function with the same signature as 'v'
	var z func(int) int

	// Assign an anonymous function to 'v'. This function takes an integer 'x' as input,
	// prints the sum of 'x' and 'x', then returns the square of 'x'.
	v = func(x int) string {
		var sum string
		sum = fmt.Sprintf("adding inside v %d", x+x)
		fmt.Println("inside v", x+x) // Print the sum of 'x' and 'x'
		return sum                   // Return the square of 'x'
	}

	// Assign the function 'y' to the variable 'z'. Now 'z' points to the function 'y'.
	z = y

	// Print the types of the variables 'v' and 'z' to confirm they are both functions that take an integer and return an integer.
	fmt.Printf("type of v: %T\n", v)
	fmt.Printf("type of z: %T\n", z)

	// Call the function stored in 'z' (which is the function 'y') with the argument 5.
	// This will return 5 * 5 = 25, and we store that value in the variable 'resulty'.
	var resulty int
	resulty = z(5) // Call 'y(5)' and store the result in 'resulty'

	// Print the result of calling 'y(5)' stored in 'resulty'.
	fmt.Println("result of y =", resulty)

	// Call the function stored in 'v' with the argument 5.
	// This will return 5 * 5 = 25, and we store that value in the variable 'result'.
	result := v(5) // Call the anonymous function stored in 'v' with the argument 5 and store the result in 'result'

	// Print the result of calling the function stored in 'v' with the argument 5.
	fmt.Println("result =", result)
}

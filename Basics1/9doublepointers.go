package main

import "fmt"

func main() {
	// Step 1: Declare an integer variable
	var number int = 10

	// Step 2: Create a pointer to store the address of `number`
	var ptr *int = &number

	// Step 3: Create a double pointer that stores the address of `ptr`
	var doublePtr **int = &ptr

	// Step 4: Print initial values
	fmt.Println("Value of number:", number)

	fmt.Println("Address of number:", &number)
	fmt.Println("\n")

	fmt.Println("Value stored in ptr (Address of number):", ptr)
	fmt.Println("Dereferenced value (*ptr):", *ptr)

	fmt.Println("Address of firstpointer:", &ptr)
	fmt.Println("\n")

	fmt.Println("Value stored in doublePtr i.e printing doubleptr (Address of ptr):", doublePtr)
	fmt.Println("Dereferenced value (*doublePtr) i.e value of stored in ptr(address of number variable):", *doublePtr)
	fmt.Println("Dereferenced twice (**doublePtr):", **doublePtr)
	fmt.Println("\n")

	// Step 5: Modify value using double pointer
	**doublePtr = 40

	// Step 6: Print updated values
	fmt.Println("\nAfter modification using double pointer i.e updating number to 40:")
	fmt.Println("Value of number:", number)
	fmt.Println("Dereferenced twice (**doublePtr):", **doublePtr)
}

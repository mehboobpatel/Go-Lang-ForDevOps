package main

import "fmt"

// Define a struct with an additional field `height`
type Person struct {
	name   string
	age    int
	height int // New field
}

// Function to modify `age` using a pointer
func updateAge(p *Person, newAge int) {
	p.age = newAge
}

// Function to modify `height` using a pointer
func updateHeight(p *Person, newHeight int) {
	p.height = newHeight
}

func main() {
	// Creating a new Person instance
	p1 := Person{"Alice", 25, 170} // Height added
	// var x *Person = &p1
	x := &p1
	fmt.Printf("Memory Address of whole struct: %p\n", x)
	fmt.Printf("Memory Address of whole struct: %p\n", &p1)
	fmt.Println("Memory Address of whole struct:", &p1)

	fmt.Println("Memory Address of Age:", &p1.age, "\n")
	fmt.Println("Memory Address of Height:", &p1.height, "\n")
	fmt.Println("Memory Address of Name:", &p1.name, "\n")

	// Display initial values

	fmt.Println("Before - Age:", p1.age, "Height:", p1.height)

	// Updating age and height using pointer functions
	updateAge(&p1, 30)
	updateHeight(&p1, 175)

	// Display updated values
	fmt.Println("After - Age:", p1.age, "Height:", p1.height)
}

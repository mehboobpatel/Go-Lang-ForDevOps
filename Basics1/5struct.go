package main

import "fmt"

// Define a struct
type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

func main() {
	// Creating an instance of Employee
	emp := Employee{
		ID:         101,
		Name:       "John Doe",
		Age:        30,
		Department: "IT",
	}

	// Accessing struct fields
	fmt.Println("Employee Name:", emp.Name)
	fmt.Println("Employee Age:", emp.Age)
}

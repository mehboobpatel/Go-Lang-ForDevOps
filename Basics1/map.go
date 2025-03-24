package main

import "fmt"

func main() {
	// Creating and initializing a map
	myMap := make(map[string]int) // Creates an empty map using make()
	fmt.Println("map of myMap", myMap)
	fmt.Println("\n")

	employeeSalary := map[string]float64{
		"Alice": 50000.50,
		"Bob":   60000.75,
		"Eve":   55000.25,
	}
	fmt.Println("INITIALLY MAP IS \n", employeeSalary)
	fmt.Println("\n")

	// Accessing map values
	fmt.Println("Alice's Salary:", employeeSalary["Alice"])
	fmt.Println("\n")

	// Adding a new key-value pair
	employeeSalary["Charlie"] = 70000.00
	fmt.Println("After Addition of Charlie\n", employeeSalary)
	fmt.Println("\n")

	// Deleting a key
	delete(employeeSalary, "Eve")
	fmt.Println("After Deletion of Eve\n", employeeSalary)

	// Checking if a key exists
	salary, exists := employeeSalary["Bob"]
	if exists {

		fmt.Println("Bob's Salary:", salary, "\n")
	} else {
		fmt.Println("Bob's Salary not found")
	}

	// Iterating over a map
	for name, salary := range employeeSalary {
		fmt.Println(name, "earns", salary)
	}
}

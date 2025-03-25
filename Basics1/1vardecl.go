package main

import "fmt"

func main() {
	// 1. Declaring a variable using the `var` keyword (no initialization)
	var name string
	var rollno int
	name = "Maheboob"
	rollno = 20
	var isVeteran bool

	fmt.Println("1. Name:", name, "rollno", rollno, "is veteran", isVeteran)
	fmt.Printf("1.1 Type of isVeteran: %T\n", isVeteran) //To find type

	// 2. Declaring and initializing a variable using the `var` keyword
	var age int = 25
	fmt.Println("2. Age:", age)

	// 3. Declaring multiple variables at once using the `var` keyword
	var (
		city    string = "Mumbai"
		pinCode int    = 400001
	)
	fmt.Println("3. City:", city, "Pin Code:", pinCode)

	// 4. Type inference (let Go infer the type based on the value)
	var country = "India"       // Go infers `country` as a string
	var population = 1380000000 // Go infers `population` as an int
	fmt.Println("4. Country:", country, "Population:", population)

	// 5. Short-hand declaration (using `:=`)
	language := "Go" // Go infers `language` as a string
	year := 2023     // Go infers `year` as an int
	fmt.Println("5. Language:", language, "Year:", year)

	// 6. Short-hand for multiple variables
	firstName, lastName := "Maheboob", "Patel"
	fmt.Println("6. Full Name:", firstName, lastName)

	// 7. Mixing types in short-hand declaration
	score, subject := 95, "Math"
	isIndian := true //Boolean
	fmt.Println("7. Score:", score, "Subject:", subject)
	fmt.Println("8. is Indian", isIndian)

}

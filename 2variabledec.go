package main

import "fmt"

func main() {
	var name string = "Maheboob" // Explicit type declaration
	age := 25                    // Implicit type inference (Go detects it's an int)
	height := 189.88             // Implicit type inference (Go detects it's an float64)
	var isDevOps bool = true     // Boolean variable
	isTest := true
	fmt.Println("Is DevOps test:", isTest)
	fmt.Printf("Type of istest: %T\n", isTest)
	var expage int = 25            // Explicitly declared as int
	var expheight float64 = 189.88 // Explicitly declared as float64
	var expweight float32 = 75.5   // Explicitly declared as float32

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("height:", height)
	fmt.Println("Is DevOps Engineer:", isDevOps)

	fmt.Println("explictage:", expage)
	fmt.Println("explictageheight:", expheight)
	fmt.Println("explictageweight:", expweight)

}

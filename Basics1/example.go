package main

import "fmt"

// If you’re working on a scientific application that requires high precision, you might use float64 and complex128:
func main() {
	// High-precision calculations
	var pi float64 = 3.14159265358979323846
	var gravity float64 = 9.80665

	// Complex number for electrical engineering
	var impedance complex128 = complex(50, 10) // 50 + 10i ohms

	fmt.Println("Pi:", pi)
	fmt.Println("Gravity:", gravity)
	fmt.Println("Impedance:", impedance)
}

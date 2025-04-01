package main

import "fmt"

// Define the Shape interface
type Shape interface {
	Area() float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Square struct
type Square struct {
	width, height float64
}

// Triangle struct
type Triangle struct {
	base, height float64
}

// Circle implements the Shape interface
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c Circle) Diameter() float64 {
	return 2 * c.radius
}

// Square implements the Shape interface
func (s Square) Area() float64 {
	return s.width * s.height
}

// Triangle implements the Shape interface
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// Function to print the area of any shape
func printArea(s Shape) {
	fmt.Println("Area of Shape:", s.Area())
}

func main() {
	// Create instances of different shapes
	var area Shape
	area = Circle{5.00} // Circle with radius 5.00
	fmt.Println("Area of Circle:", area.Area())

	square := Square{4.00, 5.00}
	triangle := Triangle{base: 6.0, height: 8.0}
	fmt.Println("Area of Circle:", Circle{(5.00)}.Area())

	fmt.Println("Area of Circle:", Circle{(5.00)}.Area())
	fmt.Println("Diameter of Circle:", Circle{(9.00)}.Diameter())

	// Print areas of different shapes using a single printArea function
	printArea(square)
	printArea(triangle)
}

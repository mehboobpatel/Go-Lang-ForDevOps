package main

import "fmt"

func main() {
	// ----------------------------------------
	// Static Arrays
	// ----------------------------------------
	// A static array has a fixed size that cannot be changed after declaration.
	// Syntax: var arrayName [size]type

	// Declare & Intiliaze a static array of integers with size 5
	var staticArray [5]int = [5]int{10, 20, 30, 40, 50}

	// Access and print array elements
	fmt.Println("Static Array:")
	for i := 0; i < len(staticArray); i++ {
		fmt.Printf("staticArray[%d] = %d\n", i, staticArray[i])
	}
	//OR JUST print directly
	fmt.Println("Static Array:", staticArray)

	// ----------------------------------------
	// Dynamic Arrays (Slices)
	// ----------------------------------------
	// A slice is a dynamically-sized array. It can grow or shrink as needed.
	// Syntax: var sliceName []type

	// Declare & Initialize a slice of integers (no size specified)
	var dynamicSlice []int = []int{1, 2, 3, 4, 5}
	// 	 dynamicSlice :=[]int{1, 2, 3, 4, 5}

	// Print the slice
	fmt.Println("\nDynamic Slice (Initial):", dynamicSlice)

	// Add elements to the slice using append()
	dynamicSlice = append(dynamicSlice, 6, 7, 8)
	fmt.Println("Dynamic Slice (After Append):", dynamicSlice)

	// Access and print slice elements
	fmt.Println("\nDynamic Slice Elements:")
	for index, value := range dynamicSlice {
		fmt.Printf("dynamicSlice[%d] = %d\n", index, value)
	}

	// ----------------------------------------
	// Key Differences Between Static Arrays and Slices
	// ----------------------------------------
	// 1. Static arrays have a fixed size, while slices are dynamic. SLICES ARE LIKE VECTOR IN CPP
	// 2. Slices are more flexible and commonly used in Go.
	// 3. Slices are reference types, while arrays are value types.

	// Example: Modifying a slice affects the underlying array
	subSlice := dynamicSlice[2:5] // Create a sub-slice from index 2 to 4
	fmt.Println("\nSub-Slice (Before Modification):", subSlice)

	// Modify the sub-slice
	subSlice[0] = 100
	fmt.Println("Sub-Slice (After Modification):", subSlice)
	fmt.Println("Dynamic Slice (After Sub-Slice Modification):", dynamicSlice)

	// ----------------------------------------
	// Common Slice Operations
	// ----------------------------------------
	// 1. Length of a slice: len(slice)
	// ACTUAL NO. OF ELEMENTS EXISITING
	fmt.Println("\nLength of Dynamic Slice:", len(dynamicSlice))

	// 2. Capacity of a slice: cap(slice)
	// ACTUAL NO. OF ELEMENTS CAPACITY THAT CAN BE STORED i.e ( Go automaticlly increase the length like Vector)
	fmt.Println("Capacity of Dynamic Slice:", cap(dynamicSlice))

	// 3. Create a slice with make()
	newSlice := make([]int, 3, 3) // Length = 3, Capacity = 5 you can use 5 aswell
	fmt.Println("\nNew Slice (Created with make):", newSlice)
	fmt.Println("Length of New Slice:", len(newSlice))
	fmt.Println("Capacity of New Slice:", cap(newSlice))
}

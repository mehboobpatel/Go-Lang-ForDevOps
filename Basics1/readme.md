# Go Programming Basics

Welcome to the Go Programming Basics guide! This guide is designed for beginners to help you get started with Go programming. Below, you'll find a comprehensive guide with topics, brief explanations, and references to the code files.

## Table of Contents
1. [Variable Declarations](#variable-declarations)
2. [Arrays and Slices](#arrays-and-slices)
3. [Structs](#structs)
4. [Maps](#maps)
5. [Pointers](#pointers)
6. [Double Pointers](#double-pointers)
7. [Pointer to Struct](#pointer-to-struct)
8. [Ranges](#ranges)
9. [Channels](#channels)

## Introduction to Go
Go, also known as Golang, is an open-source programming language developed by Google. It is designed for simplicity, efficiency, and reliability.

## Variable Declarations
Variables in Go can be declared using `var`, `const`, or shorthand `:=`.  
Go is statically typed, meaning types are determined at compile-time.  

**File:** [`vardecl.go`](vardecl.go)  
![Alt text](images/image2vartypes.png)  
![Alt text](images/image1zerovalues.png)  

## Arrays and Slices
Arrays have a fixed size, while slices are dynamic and can grow or shrink.  
Slices internally reference an array and provide flexibility in data manipulation.  

**File:** [`arrayslices.go`](arrayslices.go)  

## Structs
Structs in Go are composite data types that allow grouping multiple fields.  
They provide a way to define complex objects with different data types.  

**File:** [`struct.go`](struct.go)  
![Alt text](images/image3struct.png)  

## Maps
Maps are key-value pairs that allow fast lookups and modifications.  
Keys in a Go map must be unique, and their type must be comparable.  

**File:** [`map.go`](map.go)  
![Alt text](images/image4map.png)  
![Alt text](images/image5map.png)  

## Pointers
Pointers store memory addresses of variables, allowing direct modification of values.  
Using pointers helps in optimizing memory usage and avoiding unnecessary copies.  

**File:** [`pointers.go`](pointers.go)  

## Double Pointers
A double pointer (`**`) stores the address of another pointer.  
It is useful when you need to modify a pointer’s target indirectly.  

**File:** [`doublepointers.go`](doublepointers.go)  
![Alt text](images/image6pointers.png)  

## Pointer to Struct
Pointers can be used with structs to pass large data efficiently.  
They allow modifying struct fields without making a copy of the struct.  

**File:** [`pointrstruct.go`](pointrstruct.go)  

## Ranges
Ranges in Go are used to iterate over arrays, slices, maps, and channels.  
They return both the index and value while looping over elements.  

**File:** [`ranges.go`](ranges.go)  

## Channels
Channels in Go are used to communicate between goroutines.  
They allow you to send and receive values between different parts of your program running concurrently.

### Components of a Channel
- **Declaration:** `make(chan Type)` creates a new channel of the specified type.
- **Sending:** `channel <- value` sends a value into the channel.
- **Receiving:** `value := <-channel` receives a value from the channel.

### Example
In the provided code, two goroutines send messages into the `messages` channel. The main function waits for these messages and prints them.

**File:** [`channels.go`](channels.go)

Happy coding with Go!

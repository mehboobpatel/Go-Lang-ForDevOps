# Go Programming Basics - Advanced Topics

Welcome to the Go Programming Basics guide! This guide is designed for beginners to help you get started with advanced Go programming concepts. Below, you'll find a comprehensive guide with topics, brief explanations, and references to the code files.

---

## Table of Contents
1. [Functions as Variables](#functions-as-variables)
2. [Interfaces](#interfaces)

---

## Functions as Variables

In Go, functions are treated as first-class citizens. This means you can:
- Assign functions to variables.
- Pass functions as arguments to other functions.
- Return functions from other functions.

### What are Functions as Variables?

A **function as a variable** means you can store a function in a variable and call it later using that variable. This is similar to how you store values in variables, but instead of storing data, you store a function.

### Explanation of the Code

The code in [`1func_as_var.go`](1func_as_var.go) demonstrates:
1. Assigning a named function to a variable.
2. Assigning an anonymous function to a variable.
3. Calling functions via variables.

When you run the code, the output will demonstrate how functions can be dynamically assigned and called using variables.

---

## Interfaces

### What is an Interface?

An **interface** in Go is like a contract. It defines a set of methods that a type must have to "implement" the interface. If a type implements all the methods in an interface, it is said to satisfy the interface.

### Why Use Interfaces?

✅ **Flexibility**: Works with different types that follow the contract.  
✅ **Decoupling**: Code is not dependent on specific types.  
✅ **Reusable Code**: Functions can accept any type that implements the interface.

### Analogy:

Think of a real-world scenario: You are building a graphics application that works with various shapes (e.g., Circle, Square, Rectangle, Triangle). If you don't use interfaces, you would have to write a separate function for calculating the area for every new shape, and the rest of your code would become increasingly difficult to scale and maintain. With interfaces, you only need to create a struct for new shape type and implement make a func of that newly created struct implementing the Area() method for it. Your existing code will continue to work because it interacts with the interface, not specific types.

Think of an interface as a universal charger. For example:
- A "DeviceCharger" interface says, "You must have a `Charge()` method."
- Any device (like an iPhone or Samsung) that has a `Charge()` method can use the universal charger.
- The interface doesn't care *how* the device charges, as long as it fulfills the requirement.

### Prerequisites:
1. **Structs**: You should know how to define and use structs in Go.
2. **Methods**: You should know how to define methods for structs.

A struct does not "belong" to an interface. Instead, a struct becomes part of an interface only when it implements all the required methods.

---

### Explanation of the Code

The code in [`3interfacebasic.go`](3interfacebasic.go) demonstrates the following:

1. **Defining an Interface**:
   - The `DeviceCharger` interface defines a single method: `Charge()`.
   - Any type that has a `Charge()` method satisfies the `DeviceCharger` interface.

2. **Creating Structs**:
   - Two structs, `iPhone` and `Samsung`, are created to represent different devices.

3. **Implementing the Interface**:
   - Both `iPhone` and `Samsung` implement the `Charge()` method:
     - `iPhone`'s `Charge()` returns `"🔋 iPhone is charging with the universal charger..."`.
     - `Samsung`'s `Charge()` returns `"🔋 Samsung is charging with the universal charger..."`.

4. **Using the Interface**:
   - A variable of type `DeviceCharger` is declared.
   - An `iPhone` is assigned to the variable, and its `Charge()` method is called.
   - A `Samsung` is then assigned to the same variable, and its `Charge()` method is called.

---

### Flow of Execution:
1. The program defines a "contract" (interface) called `DeviceCharger`.
2. The `iPhone` and `Samsung` structs "apply for the job" by implementing the `Charge()` method.
3. The program assigns an `iPhone` to the `DeviceCharger` variable and calls its `Charge()` method.
4. The program assigns a `Samsung` to the same variable and calls its `Charge()` method.

---

### Output:
When you run the code  the output will be: 

for 3interfacebasic
![`3interfacebasic.go`](images/3interfacebasics.png)


for 2interfacebasic

![`2interfacebasic.go`](images/2interface.png)

for 4twomethods

![`4twomethods.go`](images/4twomethodsinterface.png)

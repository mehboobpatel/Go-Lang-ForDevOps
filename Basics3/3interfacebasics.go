package main

import "fmt"

// Step 1: Define an interface representing any device that can be charged
// 💡 DeviceCharger is like a **universal charger plug** that can charge any device
type DeviceCharger interface {
	Charge() string // Any type that has this method is considered a DeviceCharger
}

// Step 2: Create specific device structs (representing different phone brands)
// 📱 Represents an iPhone
type iPhone struct{}

// 📱 Represents a Samsung phone
type Samsung struct{}

// Step 3: Implement the interface method for iPhone
// 📌 This function makes `iPhone` compatible with `DeviceCharger`
func (phone iPhone) Charge() string {
	// phone is a input variable of type iPhone
	// `phone` is a receiver of type `iPhone`, meaning this function belongs to the `iPhone` struct
	// This function implements the `Charge` method, making `iPhone` part of the `DeviceCharger` interface
	return "🔋 iPhone is charging with the universal charger..."
}

// Step 4: Implement the interface method for Samsung
// 📌 This function makes `Samsung` compatible with `DeviceCharger`
func (phone Samsung) Charge() string {
	// phone is a input variable of type Samsung

	// `phone` is a receiver of type `Samsung`, meaning this function belongs to the `Samsung` struct
	// This function implements the `Charge` method, making `Samsung` part of the `DeviceCharger` interface
	return "🔋 Samsung phone is charging with the universal charger..."
}

// Step 5: Use the interface in the main function
func main() {
	// Step 6: Declare a variable of type DeviceCharger (a universal charger)
	// 💡 `myCharger` is like a universal charging port — it can charge **any device** that has the `Charge()` method
	var myCharger DeviceCharger

	// Step 7: Assign an iPhone to the charger

	//iPhone is a struct type (a user-defined data type).
	//iPhone{} is a struct value (an instance of the struct).
	//There is no function call here, just the creation of a new iPhone object.
	myCharger = iPhone{}            // 🔌 Plugging in an iPhone
	fmt.Println(myCharger.Charge()) // Output: 🔋 iPhone is charging with the universal charger...

	// Step 8: Assign a Samsung phone to the charger
	myCharger = Samsung{}           // 🔌 Plugging in a Samsung phone
	fmt.Println(myCharger.Charge()) // Output: 🔋 Samsung phone is charging with the universal charger...
}

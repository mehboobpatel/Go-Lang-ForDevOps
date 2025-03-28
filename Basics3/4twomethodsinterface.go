package main

import "fmt"

// Step 1: Define an interface representing any device that can be charged
type DeviceCharger interface {
	Charge() string  // Any type with this method fits into DeviceCharger
	Storage() string // This method is missing in iPhone
}

// Step 2: Create a struct representing an iPhone with fields
type iPhone struct {
	Model     string // Model of the iPhone
	Battery   int    // Battery percentage
	StorageGB int    // Storage capacity in GB
}

// Step 3: Implement the Charge() method for iPhone
func (phone iPhone) Charge() string {
	fmt.Println(phone.Model, "Charging... Battery:", phone.Battery, "%") // Print the charging message
	return fmt.Sprintf("🔋 %s is charging... Battery: %d%%", phone.Model, phone.Battery)
}

// Step 4: Implement the Storage() method for iPhone
func (phone iPhone) Storage() string {
	return fmt.Sprintf("💾 %s has %dGB storage available.", phone.Model, phone.StorageGB)
}

// Step 5: Use the interface in the main function
func main() {
	// Declare a variable of type DeviceCharger (a universal charger)
	var myCharger DeviceCharger

	// Assign an iPhone instance (now it has both Charge() and Storage())
	myCharger = iPhone{Model: "iPhone 15 Pro", Battery: 50, StorageGB: 256}

	// Call the Charge() and Storage() methods
	fmt.Println(myCharger.Charge())  // Output: 🔋 iPhone 15 Pro is charging... Battery: 50%
	fmt.Println(myCharger.Storage()) // Output: 💾 iPhone 15 Pro has 256GB storage available.

	// Assign another iPhone instance
	myCharger = iPhone{Model: "Samsung Note", Battery: 90, StorageGB: 128}
	fmt.Println(myCharger.Charge())  // Output: 🔋 Samsung Note is charging... Battery: 90%
	fmt.Println(myCharger.Storage()) // Output: 💾 Samsung Note has 128GB storage available.
}

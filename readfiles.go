package main

import (
	"fmt"
	"os"
)

func getmehboob() {
	fmt.Println("calling function outside main")
}
func main() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File Contents:\n", string(data))

	getmehboob()
}

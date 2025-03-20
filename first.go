package main
import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <your_name>")
		return
	}
	fmt.Println("Hello,", os.Args[1])
}

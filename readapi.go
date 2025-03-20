package main
import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api64.ipify.org?format=text")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	ip, _ := io.ReadAll(resp.Body)
	fmt.Println("Your Public IP:", string(ip))
}

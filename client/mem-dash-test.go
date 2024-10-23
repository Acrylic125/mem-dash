package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	_, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to reach port 6379")
		os.Exit(1)
	}
	fmt.Println("Hello, World!")
}

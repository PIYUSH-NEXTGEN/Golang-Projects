package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("OS:", runtime.GOOS)                // Gets the operating system
	fmt.Println("Architecture:", runtime.GOARCH)    // Gets the CPU architecture
	fmt.Println("CPU cores:", runtime.NumCPU())     // Gets the number of CPU cores
	fmt.Println("Go version:", runtime.Version())   // Gets the installed Go version
	fmt.Println("Username:", os.Getenv("USERNAME")) // Gets the current username

	hostname, err := os.Hostname() // Gets the computer's hostname

	if err != nil {
		fmt.Println("Error:", err) // Prints the error if hostname lookup fails
		return                     // Stops the program
	}

	fmt.Println("Hostname:", hostname) // Displays the hostname
}

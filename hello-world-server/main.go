package main

import (
	"fmt"
	"net/http" // Provides tools for creating HTTP servers, handling requests, and sending responses
)

// helloHandler handles HTTP requests sent to the "/" route.
// w is the ResponseWriter used to send data back to the client/browser.
// r is a pointer to the Request containing information about the incoming HTTP request.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, there") // Writes "Hello, there" into the HTTP response sent to the browser
}

func main() {
	// Registers helloHandler to handle requests made to the "/" URL path.
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server running on http://localhost:8080")

	// Starts the HTTP server on port 8080.
	// The second argument nil tells Go to use its default HTTP request multiplexer (router).
	err := http.ListenAndServe(":8080", nil)

	// ListenAndServe returns an error if the server cannot start or stops unexpectedly.
	if err != nil {
		fmt.Println(err) // Prints the error to the terminal
	}
}

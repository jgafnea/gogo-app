package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World")
}

func main() {
	// Register the handler function for the root ("/") route
	http.HandleFunc("/", helloHandler)

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

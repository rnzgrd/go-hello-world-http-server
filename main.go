package main

import (
	"fmt"
	"net/http"
)

// handler function to handle HTTP requests
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Assign the helloWorldHandler function to handle requests to the root URL "/"
	http.HandleFunc("/", helloWorldHandler)

	// Start the server on port 8080 and log any errors
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

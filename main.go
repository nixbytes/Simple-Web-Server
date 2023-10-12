package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler handles incoming HTTP requests that contain form data
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the request body
	if err := r.ParseForm(); err != nil {
		// If there's an error parsing the form, send an error response to the client
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Indicate to the client that the post request has been processed successfully
	fmt.Fprintf(w, "Post request successful")

	// Extract 'name' and 'address' fields from the form data
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Send the received 'name' and 'address' back to the client as a response
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// helloHandler responds to HTTP requests made to the "/hello" path
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the requested URL path is exactly "/hello"
	if r.URL.Path != "/hello" {
		// If not, send a "404 not found" error response
		http.Error(w, "404 not found!", http.StatusNotFound)
		return
	}

	// Check if the HTTP request method is "GET"
	if r.Method != "GET" {
		// If not, send an error response indicating the method is not supported
		http.Error(w, "Method not Supported", http.StatusNotFound)
		return
	}

	// Send "hello" as a response to the client
	fmt.Fprintf(w, "hello")
}

func main() {
	// Create a file server to serve static files from the "./static" directory
	fileserver := http.FileServer(http.Dir("./static"))

	// Set the root path ("/") to use the file server for serving static files
	http.Handle("/", fileserver)

	// Register the formHandler and helloHandler functions to their respective paths
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Print a message indicating the server is starting on port 8080
	fmt.Println("Starting Server on Port :8080")

	// Start the HTTP server on port 8080. If there's an error, log it and terminate the program.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.Handle("/form", formHandler)
	http.Handle("/hello", helloHandler)
	fmt.Println("Starting Server on Port :8080")

	if err := http.ListerAndServe; err != nil {

		log.Fatal(err)
	}
}

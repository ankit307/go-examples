package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, req *http.Request) {
	log.Println("Request received")
	fmt.Fprintf(w, "Hello World")
}

// Middleware logging function to log request
func logging(nextHandler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received : %v\n", r)
		fmt.Printf("Path:%s\n", r.URL.Path)
		fmt.Println("Method :", r.Method)
		nextHandler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", logging(helloWorld))
	log.Println("Server started at :3000")
	http.ListenAndServe(":3000", nil)
}

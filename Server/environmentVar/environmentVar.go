package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("PORT", ":3000")
	port := os.Getenv("PORT")
	fmt.Println("Loaded Environment Variable:", port)
}

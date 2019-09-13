package main

import (
	"fmt"
)

func main() {
	days := make(map[int]string)
	days[1] = "One"
	days[2] = "Two"
	days[3] = "Three"

	for key, value := range days {
		fmt.Printf("Key: %d || Value: %s\n", key, value)
	}

}

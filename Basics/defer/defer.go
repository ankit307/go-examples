package main

import (
	"fmt"
)

func worker() {
	i := []int{1, 2, 3, 4, 5}
	defer fmt.Println("Completed")
	for val := range i {
		fmt.Println(val)
	}
}

func main() {
	go worker()
	fmt.Scanln()
}

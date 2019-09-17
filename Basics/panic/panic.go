package main

import (
	"fmt"
)

func loop() {
	i := []int{1, 2, 3, 4, 5}
	for val := range i {
		fmt.Println(val)
	}
}

func main() {
	panic("Some problem occured")

	loop()
}

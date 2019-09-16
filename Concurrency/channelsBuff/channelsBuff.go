package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "String 1"
	c <- "String 2"
	fmt.Println(<-c)
	fmt.Println(<-c)

}

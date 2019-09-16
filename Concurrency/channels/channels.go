package main

import (
	"fmt"
)

func f(from string, ch chan<- string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
	ch <- "done"
}

func main() {
	c := make(chan string)
	go f("goroutine", c)
	fmt.Println(<-c)
}

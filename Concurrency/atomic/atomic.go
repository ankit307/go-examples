package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var balance uint64
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&balance, 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Balance :", balance)
}

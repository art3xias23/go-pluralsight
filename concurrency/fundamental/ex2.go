package main

import (
	"fmt"
	"sync"
)

func Example2() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		fmt.Println("this happens asynchronously")
		wg.Done()
	}()

	fmt.Println("this is synchronous")
	wg.Wait()

}

package main

import (
	"fmt"
	"sync"
)

func Example1() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		fmt.Println("do some async thing")
		wg.Done()

	}()

	wg.Wait()
}

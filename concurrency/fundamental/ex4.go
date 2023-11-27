package main

import "fmt"

func Example4() {
	ch1, ch2 := make(chan string), make(chan string)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}

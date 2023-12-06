package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	orderCh := make(chan newOrder)
	validCh := make(chan newOrder)
	invalidCh := make(chan newOrder)
	go getOrders(orderCh)
	go validateOrders(orderCh, validCh, invalidCh)
	wg.Add(1)
	go func() {
		order := <-validCh
		fmt.Printf("Valid order, %v", order.Name)
		wg.Done()
	}()
	go func() {
		order := <-invalidCh
		fmt.Printf("Invalid order, %v", order.Name)
		wg.Done()
	}()
	wg.Wait()
}

func getOrders(out chan newOrder) {
	for _, item := range jso {
		var no newOrder
		erra := json.Unmarshal([]byte(item), &no)
		if erra != nil {
			fmt.Println("ERR, %v", erra)
		}

		out <- no
	}
}

func validateOrders(inOrders, outValid, outInvalid chan newOrder) {
	order := <-inOrders

	if order.Age < 18 {
		outInvalid <- order
		return
	}
	outValid <- order
}

var jso = []string{
	`{"name":"Tino", "age":10}`,
	`{"name":"Bino", "age":20}`,
	`{"name":"Cino", "age":10}`}

type newOrder struct {
	Name string
	Age  int
}

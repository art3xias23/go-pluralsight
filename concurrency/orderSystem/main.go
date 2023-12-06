package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	exampleOne()
}

func exampleTwo() {
	var wg sync.WaitGroup
	fmt.Println("1")
	wg.Add(1)
	go someGoRoutine(&wg)
	wg.Wait()
	fmt.Println("2")
}

func someGoRoutine(wg *sync.WaitGroup) {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func validateOrders(in, out chan order, errCh chan invalidOrder) {
	fmt.Println("entered val or")
	order := <-in
	if order.Quantity <= 0 {
		errCh <- invalidOrder{order: order, err: errors.New("quantity must be greater than 0")}
	} else {
		out <- order
	}
}

func exampleOne() {

	var wg sync.WaitGroup
	var receiveOrdersCh = make(chan order)
	var validOrderCh = make(chan order)
	var inValidOrderCh = make(chan invalidOrder)
	go receiveOrders(receiveOrdersCh)
	fmt.Println("passed rec or")
	go validateOrders(receiveOrdersCh, validOrderCh, inValidOrderCh)
	fmt.Println("passed val order")
	wg.Add(1)
	go func() {
		fmt.Println("entered go 1")
		order := <-validOrderCh
		fmt.Println("Valid order received: %v\n", order)
		wg.Done()
	}()
	go func() {
		fmt.Println("entered go 2")
		order := <-inValidOrderCh
		fmt.Println("Invalid order received: %v\n", order)
		wg.Done()
	}()
	wg.Wait()
}

func receiveOrders(out chan order) {
	fmt.Println("entered rec or")
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		out <- newOrder
	}
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}

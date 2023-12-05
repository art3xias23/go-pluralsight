package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func main() {
	exampleTwo()
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

func exampleOne() {

	var wg sync.WaitGroup
	wg.Add(1)
	go receiveOrders(&wg)
	wg.Wait()
	fmt.Printf("%v", orders)
}

func receiveOrders(wg *sync.WaitGroup) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		orders = append(orders, newOrder)
	}
	wg.Done()
}

var orders []order
var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}

package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed numbers.txt
	data []byte
)

func main() {

	items := strings.Split(string(data), "\n")
	var sum int
	for _, line := range items {
		if line != "" {
			number, _ := strconv.Atoi(line)
			sum += number
		}
	}

	fmt.Printf("Total: %v", sum)
}

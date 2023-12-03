package main

import (
	"bytes"
	"fmt"
	"strings"
)

type printer interface {
	Print() string
}

type user struct {
	username string
	id       int
}

func (u user) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}

type menuItem struct {
	name   string
	prices map[string]float64
}

func (mi menuItem) Print() string {
	var buffer bytes.Buffer
	buffer.WriteString(mi.name)
	buffer.WriteString(strings.Repeat("-", 10) + "\n")
	for size, price := range mi.prices {
		fmt.Fprintf(&buffer, "\t%10s%10.2f\n", size, price)
	}

	return buffer.String()
}

func main() {
	// printSomeValues()
	// printVarAndFunc()
	printSlice()
}

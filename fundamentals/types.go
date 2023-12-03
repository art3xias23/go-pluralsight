package main

import (
	"fmt"
)

func printVars() {

	//Reassigning creates a new value
	var a = 10
	var b = a

	fmt.Printf("a:%p", &a)
	fmt.Printf("b:%p", &b)
}
func printFunc() {

	var a = 10
	fmt.Printf("a:%p", &a)
	//Passing in a function creates a new value
	PassByValue(a)
}

func PassByValue(a int) {

	fmt.Printf("a in func:%p", &a)
}

func printSlice() {
	a := []float64{
		97.5,
		105,
		63.5}

	b := a

	c := a
	fmt.Printf("%p", &a[0])
	fmt.Println("\n")
	fmt.Printf("%p", &b[0])
	fmt.Println("\n")
	fmt.Printf("%p", &c[0])

}

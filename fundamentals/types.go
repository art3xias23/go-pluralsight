package main

import (
	"fmt"
)

func printVars() {

	//Reassigning creates a new value
	var a = 10
	var b = a

	fmt.Printf("a:%p\nb:%p", &a, &b)
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

	fmt.Printf("The variables have different memory allocations")
	fmt.Print("\n")
	fmt.Printf("a: %p", &a)
	fmt.Print("\n")
	fmt.Printf("b: %p", &b)
	fmt.Print("\n")
	fmt.Print("------------------")
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Printf("The underlying structure keeps it's reference")
	fmt.Print("\n")
	fmt.Printf("a[0]: %p", &a[0])
	fmt.Print("\n")
	fmt.Printf("b[0]: %p", &b[0])
	fmt.Print("\n")
	fmt.Print("------------------")
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Printf("Confirming they hold the same values")
	fmt.Print("\n")
	fmt.Printf("a: %v", a)
	fmt.Print("\n")
	fmt.Printf("b: %v", b)

}

func passByReference() {
	var a = 5

	fmt.Printf("a: %p", &a)
	fmt.Print("\n")
	passByReferenceFunc(&a)
}
func passByReferenceFunc(a *int) {
	fmt.Printf("f: %p", a) // dereference value
}

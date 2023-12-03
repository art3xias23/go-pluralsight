package main

import "fmt"

func main() {
	initSlice := []float32{
		12.3,
		40.1,
		57.9}

	referenceSlice := initSlice

	fmt.Println(&initSlice[0])
	fmt.Println(&referenceSlice[0])

	valueSlice := copy(initSlice)
	fmt.Println(&valueSlice[0])
}

func copy[V any](s []V) []V {
	copy := make([]V, len(s))

	for index, value := range s {
		copy[index] = value
	}

	return copy
}

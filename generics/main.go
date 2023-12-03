package main

import "fmt"

func main() {
	initSlice := []float64{
		12.3,
		40.1,
		57.9}

	referenceSlice := initSlice

	fmt.Println(&initSlice[0])
	fmt.Println(&referenceSlice[0])

	valueSlice := copy(initSlice)
	fmt.Println(&valueSlice[0])
}

func copy(s []float64) []float64 {
	copy := make([]float64, len(s))

	for index, value := range s {
		copy[index] = value
	}

	return copy
}

package main

import "fmt"

func main() {
	// maps()
	genConcat()
}

func slices() {

	initSlice := []float32{
		12.3,
		40.1,
		57.9}

	referenceSlice := initSlice

	fmt.Println(&initSlice[0])
	fmt.Println(&referenceSlice[0])

	valueSlice := copySlice(initSlice)
	fmt.Println(&valueSlice[0])
}

func copySlice[V any](s []V) []V {
	copy := make([]V, len(s))

	for index, value := range s {
		copy[index] = value
	}

	return copy
}

func maps() {
	myMap := map[float32]string{1: "Ben",
		2: "Michael",
		3: "Brian"}
	copy := copyMaps(myMap)

	fmt.Print(copy)

}

func copyMaps[K comparable, V any](m map[K]V) map[K]V {
	copy := make(map[K]V, len(m))

	for key, value := range m {
		copy[key] = value
	}

	return copy
}

type addable interface {
	int | float32 | string
}

func genConcat() {
	integers := []int{1, 2, 3}
	floats := []float32{1.3, 2.3, 3.4}
	strs := []string{"I", "want", "ice-cream"}
	result := concat(integers)
	fmt.Println(result)

	result2 := concat(floats)
	fmt.Println(result2)

	result3 := concat(strs)
	fmt.Println(result3)
}

func concat[V addable](s []V) V {
	var result V
	for _, val := range s {
		result += val
	}
	return result
}

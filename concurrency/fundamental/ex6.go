package main

import "fmt"

func Example6() {
	var before int = 42

	fmt.Printf("before: %d\n", before)

	changeNum(&before)

	fmt.Printf("after: %d\n", before)

}

func changeNum(num *int) {
	fmt.Printf("num address: %d\n", num)
	*num = 6
	fmt.Printf("num value: %d\n", *num)
}

package main

import "fmt"

func main() {
	nums := []float32{1.000000, 2, 3, 4, 5}

	for index, value := range nums {
		fmt.Printf("value at index %x is %v\n", index, value)
	}
}

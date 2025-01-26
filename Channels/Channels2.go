package main

import (
	"fmt"
)

func process_numChan(num chan int) {

	for value := range num {
		fmt.Println("this is the value", value)
	}

}

func main() {

	numChan := make(chan int)

	go process_numChan(numChan)

	numChan <- 8

	numChan <- 9

}

package main

import (
	"fmt"
)

func process_numChan(num chan int) {
	for value := range num {
		fmt.Println("this is the cuurent value in the chhannel", value)
	}
}
func main() {
	numChan := make(chan int)

	go process_numChan(numChan)
	for i := 0; i < 100; i++ {
		numChan <- i
	}

}

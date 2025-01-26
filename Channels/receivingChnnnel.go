package main

import "fmt"

func sum(result chan int, a int, b int) {
	additon := a + b
	result <- additon
}

func main() {
	result := make(chan int)

	go sum(result, 5, 6)

	received_channel_data := <-result

	fmt.Println(received_channel_data)
}

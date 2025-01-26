package main

import "fmt"

func main() {
	msgChannel := make(chan string)

	msgChannel <- "ping" //channels are blocking in nature, until the receiver is ready to receive it.

	receiver := msgChannel

	fmt.Println(receiver);
}
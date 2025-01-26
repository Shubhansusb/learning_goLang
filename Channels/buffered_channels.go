package main

import (
	"fmt"
)

func main() {
	chan_email := make(chan string, 100) //buffered channel is nonblocking till 100.

	chan_email <- "loda@gmail.com"
	chan_email <- "mc@gmail.com"

	fmt.Println(<-chan_email)
	fmt.Println(<-chan_email)
}

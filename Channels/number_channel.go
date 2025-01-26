package main

import (
	"fmt"
	"time"
)

func process_number_channel(numChan chan int) {
	fmt.Println("processing the number", <-numChan)
}

func main() {

	numberChan := make(chan int) //defining the channel

	go process_number_channel(numberChan) // pricessing the number , this will run at the same instacne when the main will be executed.
	numberChan <- 5                       //putting the vaue in the channel as 5

	time.Sleep(time.Millisecond * 100)

}

 
// Key Takeaways for Beginners
// Channels in Go Are Blocking:

// A send (channel <- value) will block until there is a receiver ready to read from the channel.
// Similarly, a receive (<-channel) will block until there is a sender sending a value into the channel.

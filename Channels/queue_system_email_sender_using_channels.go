package main

import (
	"fmt"
	"time"
)

func email_sender(queue_like_cahnnel chan string, done chan bool) {

	defer func() {
		done <- true
	}()

	for email := range queue_like_cahnnel {
		fmt.Println("sending mail to :", email)
		time.Sleep(time.Second);
	}

}
func main() {

	email_chan := make(chan string, 100) //unblocking till 100,so the loop will fill it with 100 strings without unblocking.
	sync_channel := make(chan bool)
	go email_sender(email_chan, sync_channel)

	for i := 0; i < 100; i++ {
		email_chan <- fmt.Sprintf("%d@gmail.com", i)
	}
	close(email_chan) // if we will not close the channel, the infinite loop on the quueu_like_cahnnel will go on indefinitely and defer fun will not happen and sync_cahnnel will keep on waiting leading to the deadlock.
	<-sync_channel

}

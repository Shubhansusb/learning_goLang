package main

import "fmt"

func something(done chan bool) {
	defer func(){
		done<-true
	}(); //defer will makke it run after the printLn is done, and will put something in done channel

	fmt.Println("Shubhansu singh bhadoria");
}

func main() {

	some_channel := make (chan bool);

	go something(some_channel);
  <-some_channel //this is blocking 
}


//this is alternative of using wait groups.
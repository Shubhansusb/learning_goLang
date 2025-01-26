package main

import (
	"fmt"
	"time"
)

func performTask(task int) {
	fmt.Println("performing task", task)
}
func main() {
	for i := 0; i <= 10; i++ {
		go performTask(i)
	}

	// for j := 0; j <= 100; j++ {
	// 	 go fmt.Println("Shubhansu")
	// }

	time.Sleep(time.Millisecond*2); // we are making the main fun to sleep, so that go routines may perform the exection, because as we use the gp keyword, it makes the paticualr operation non blocking.
	// for i := 0; i <= 10; i++ {
	// 		performTask(i);
	// }
}

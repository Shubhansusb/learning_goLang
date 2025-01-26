package main

import (
	"fmt"
	"sync"
)

func performTask(task int, w* sync.WaitGroup) {
	defer w.Done();
	fmt.Println("performing task", task)
}
func main() {
	var ws sync.WaitGroup 
	for i := 0; i <= 10; i++ {
		ws.Add(1);
		go performTask(i, &ws)
	}

	ws.Wait(); 
	
}

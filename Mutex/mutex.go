//mutex or mutual exclusion are used while concurrency, to prevent race condition.

package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) increment(wg *sync.WaitGroup) {

	defer func() {
		wg.Done()
		p.mu.Unlock()

	}()
	p.mu.Lock()
	p.views += 1

}
func main() {

	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.increment(&wg)
	}

	wg.Wait()
	fmt.Println("view are ", myPost.views)

}

//the above code might lead to race conditon as our loop is launchinf 100 go routines,

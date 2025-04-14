package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1

}

func main() {
	wg := sync.WaitGroup{}

	mypost := post{views: 0}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go mypost.inc(&wg)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println(mypost.views)
}

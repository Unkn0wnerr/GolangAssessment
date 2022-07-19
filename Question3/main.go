package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, mtx *sync.Mutex) {
	mtx.Lock()
	x = x + 1
	mtx.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var mtx sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &mtx)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

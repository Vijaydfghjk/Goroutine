package main

import (
	"fmt"
	"sync"
)

var (
	slice []int
	wg    sync.WaitGroup
	mu    sync.Mutex
)

func addElement(val int) {
	defer wg.Done()
	mu.Lock()
	slice = append(slice, val)
	mu.Unlock()
}

func main() {
	wg.Add(3)

	go addElement(1)
	go addElement(2)
	go addElement(3)

	wg.Wait()

	fmt.Println("Final slice:", slice)
}


1. Here without mutex in the 3 goroutine any one or two same time access the slice so what happen
   in the slice they can't add all vlue. this is problem.

2. That's what race condition is coming which is mutex.

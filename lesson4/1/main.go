package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var total = 0
	var workers = make(chan int, 1000)
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		workers <- 1

		go func(m *sync.Mutex) {
			m.Lock()
			total += <-workers
			m.Unlock()
		}(&m)
	}

	time.Sleep(time.Second * 1)
	fmt.Printf("Variable value %v\n", total)
}

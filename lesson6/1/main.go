package main

import (
	"os"
	"runtime/trace"
	"sync"
)

const count = 1000

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		arr  [count]int
		lock sync.Mutex
		wg   sync.WaitGroup
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func(i int) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			if i%2 == 0 {
				arr[i] = i
			} else {
				arr[i] = i + 1
			}
		}(i)
	}
	wg.Wait()
}

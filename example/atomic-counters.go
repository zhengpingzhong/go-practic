package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
				// 如下操作得不到50000
				//ops++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
}

// ops: 50000

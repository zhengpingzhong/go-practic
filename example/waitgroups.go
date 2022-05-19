package main

import (
	"fmt"
	"sync"
	"time"
)

func workerWaitgroups(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			workerWaitgroups(i)
		}()
	}
	wg.Wait()
}

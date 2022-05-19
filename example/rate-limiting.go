package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

//request 1 2022-05-11 11:13:27.9126231 +0800 CST m=+0.208728501
//request 2 2022-05-11 11:13:28.1169894 +0800 CST m=+0.413094801
//request 3 2022-05-11 11:13:28.320788 +0800 CST m=+0.616893401
//request 4 2022-05-11 11:13:28.5108056 +0800 CST m=+0.806911001
//request 5 2022-05-11 11:13:28.7131869 +0800 CST m=+1.009292301
//request 1 2022-05-11 11:13:28.7131869 +0800 CST m=+1.009292301
//request 2 2022-05-11 11:13:28.7131869 +0800 CST m=+1.009292301
//request 3 2022-05-11 11:13:28.7131869 +0800 CST m=+1.009292301
//request 4 2022-05-11 11:13:28.9141219 +0800 CST m=+1.210227301
//request 5 2022-05-11 11:13:29.114928 +0800 CST m=+1.411033401

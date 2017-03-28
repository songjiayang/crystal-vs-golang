package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	n := 100000
	channel := make(chan int, 1)
	for i := 0; i < n; i++ {
		go func() {
			time.Sleep(1 * time.Millisecond) // # io block 1ms
			channel <- 1
		}()
	}

	total := 0
	for i := 0; i < n; i++ {
		total += <-channel
	}

	close(channel)
}

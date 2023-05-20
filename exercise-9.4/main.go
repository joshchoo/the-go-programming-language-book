package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	var numPipelines int64 = 1000000
	startCh := make(chan int)
	var endCh <-chan int = startCh

	log.Println("Creating pipeline...")
	var i int64
	for ; i < numPipelines; i++ {
		endCh = makePipeline(endCh)
	}
	log.Printf("Created pipeline of %d goroutines.\n", i)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Total stack in use: %d bytes\n", m.StackInuse)
	log.Printf("Stack use per goroutine: %d bytes\n", m.StackInuse/uint64(i))

	go func() {
		log.Println("Sending 0")
		startCh <- 0
	}()
	t := time.Now()

	n := <-endCh
	close(startCh)

	sendDuration := time.Since(t).Milliseconds()
	log.Printf("Done. Received value: %d. Took %d ms\n", n, sendDuration)
	log.Printf("Send+Receives per ms: %d", 2*i/sendDuration)

	for range endCh {
		// Wait until endCh is closed. This occurs when all goroutines have exited cleanly.
	}
	log.Println("Done.")
}

func makePipeline(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 1
		}
		close(out)
	}()
	return out
}

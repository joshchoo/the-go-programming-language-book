package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	pingCh := make(chan struct{})
	pongCh := make(chan struct{})

	go func() {
		wg.Add(1)
		var count int64 = 0
		started := time.Now()

		for range pingCh {
			// Increment by two because a Ping and Pong occur everytime count is incremented.
			count += 2
			if count%1000000 == 0 {
				ratePerMs := count / time.Since(started).Milliseconds()
				log.Printf("%d communications per ms\n", ratePerMs)
			}

			pongCh <- struct{}{}
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for range pongCh {
			pingCh <- struct{}{}
		}
		wg.Done()
	}()

	pingCh <- struct{}{}
	log.Println("Started")
	wg.Wait()
}

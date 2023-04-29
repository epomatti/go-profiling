package main

import (
	"sync"
	"time"
)

func blockingDemo4() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var m sync.Mutex
		m.Lock()
		time.Sleep(3 * time.Second)
		m.Unlock()
	}()
	wg.Wait()
}

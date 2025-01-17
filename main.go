package main

import (
	"sync"
	"time"
)

func main() {
	cmd := NewCmdFlags()
	routes := loadRequest(cmd.FileDescriptor)
	var wg sync.WaitGroup
	wg.Add(cmd.nbR * len(routes.Requests))
	var resp = &Rapport{mapping: make(map[string]map[string][]float32)}
	interval := waitForRate(cmd.RateLimit, cmd.Freq)
	now := time.Now()
	// Limit goroutines spawn
	sem := make(chan struct{}, cmd.c)
	for i := 1; i <= cmd.nbR; i++ {
		for _, req := range routes.Requests {
			sem <- struct{}{}
			// Freq fire request (e.g 10req/S)
			time.Sleep(interval)
			go req.fireRequest(&wg, resp, &sem)
		}
	}
	wg.Wait()
	close(sem)
	resp.display(cmd.nbR, time.Since(now))
}

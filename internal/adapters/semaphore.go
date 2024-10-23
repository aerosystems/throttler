// Package adapters provides the implementation of the APIAdapter interface.
package adapters

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// SemaphoreAPIAdapter provides the implementation of the APIAdapter interface.
type SemaphoreAPIAdapter struct {
	url       string
	rateLimit int
	client    *http.Client
}

// NewSemaphoreAPIAdapter creates a new instance of SemaphoreAPIAdapter.
func NewSemaphoreAPIAdapter(url string, rateLimit int) *SemaphoreAPIAdapter {
	return &SemaphoreAPIAdapter{
		url:       url,
		rateLimit: rateLimit,
		client:    &http.Client{},
	}
}

// Push sends a request to the API.
func (f *SemaphoreAPIAdapter) Push(count int) {
	throttler := make(chan struct{}, f.rateLimit)
	var wg sync.WaitGroup
	wg.Add(count)
	for range make([]struct{}, count) {
		throttler <- struct{}{}
		go func() {
			defer func() {
				<-throttler
				wg.Done()
			}()
			req, _ := f.client.Post(f.url, "application/json", nil)
			defer req.Body.Close() //nolint: errcheck
			fmt.Printf("Semaphore implementation status: %s\n", req.Status)
			time.Sleep(1 * time.Second)
		}()
	}
	wg.Wait()
}

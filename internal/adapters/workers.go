// Package adapters provides the implementation of the APIAdapter interface.
package adapters

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// WorkerAPIAdapter provides the implementation of the APIAdapter interface.
type WorkerAPIAdapter struct {
	url          string
	workersCount int
	client       *http.Client
}

// NewWorkerAPIAdapter creates a new instance of WorkerAPIAdapter.
func NewWorkerAPIAdapter(url string, rateLimit int) *WorkerAPIAdapter {
	return &WorkerAPIAdapter{
		url:          url,
		workersCount: rateLimit,
		client:       &http.Client{},
	}
}

// Push sends a request to the API.
func (f *WorkerAPIAdapter) Push(count int) {
	if count < f.workersCount {
		f.workersCount = count
	}

	var (
		requestIDChannel = make(chan struct{}, count)
		workerComplete   = new(sync.WaitGroup)
	)

	workerComplete.Add(f.workersCount)
	for range make([]struct{}, f.workersCount) {
		go f.workerFetchRepositoryStarByID(workerComplete, requestIDChannel)
	}

	for range make([]struct{}, count) {
		requestIDChannel <- struct{}{}
	}

	close(requestIDChannel)

	workerComplete.Wait()
}

func (f *WorkerAPIAdapter) workerFetchRepositoryStarByID(wg *sync.WaitGroup, requestIDChannel <-chan struct{}) {
	for range requestIDChannel {
		req, _ := f.client.Post(f.url, "application/json", nil)
		defer req.Body.Close() //nolint: errcheck
		fmt.Printf("Workers implementation status: %s\n", req.Status)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

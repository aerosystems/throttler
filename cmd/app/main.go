// Package main provides the entry point for the application.
package main

import (
	"github.com/aerosystems/throttler/internal/adapters"
	"github.com/aerosystems/throttler/internal/services"
)

const (
	apiURL        = "http://localhost/api-mock"
	rateLimit     = 10
	countRequests = 100
)

func main() {
	workersAPIAdapter := adapters.NewWorkerAPIAdapter(apiURL, rateLimit)
	semaphoreAdapter := adapters.NewSemaphoreAPIAdapter(apiURL, rateLimit)

	fooGenerator := services.NewQueryGenerator(workersAPIAdapter, countRequests)
	fooGenerator.Run()

	barGenerator := services.NewQueryGenerator(semaphoreAdapter, countRequests)
	barGenerator.Run()
}

// Package services provides the business logic for the application.
package services

// QueryGenerator provides the business logic for the application.
type QueryGenerator struct {
	apiAdapter    APIAdapter
	countRequests int
}

// NewQueryGenerator creates a new instance of QueryGenerator.
func NewQueryGenerator(apiAdapter APIAdapter, countRequests int) *QueryGenerator {
	return &QueryGenerator{
		apiAdapter:    apiAdapter,
		countRequests: countRequests,
	}
}

// Run sends a request to the API.
func (f *QueryGenerator) Run() {
	f.apiAdapter.Push(f.countRequests)
}

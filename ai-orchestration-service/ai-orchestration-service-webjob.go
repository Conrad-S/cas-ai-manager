package ai_orchestration_service

import (
	"fmt"
	"time"
)

// WebJob represents a periodic task that updates the cache
type WebJob struct {
	cache *Cache
}

// NewWebJob creates a new WebJob
func NewWebJob(cache *Cache) *WebJob {
	return &WebJob{
		cache: cache,
	}
}

// Start begins the periodic cache update and triggers an immediate update
func (wj *WebJob) Start() {
	// Trigger an immediate update on server start
	wj.updateCache()

	// Start the periodic update (every hour)
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			wj.updateCache()
		}
	}()
}

// Update cache function in WebJob
func (wj *WebJob) updateCache() {
	fmt.Println("Updating cache with new model endpoints...")

	// Ensure the keys match what you're searching for in modelHandler
	modelURLs := map[string]string{
		"gpt-3.5-v1.021": "https://api.openai.com/v1/gpt-3.5-v1.021",
		"gpt-4-v2.002":   "https://api.openai.com/v1/gpt-4-v2.002",
	}

	for key, url := range modelURLs {
		wj.cache.Set(key, url, 24*time.Hour)
	}
}

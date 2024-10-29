package ai_orchestration_service

import (
	"sync"
	"time"
)

// CacheEntry represents a single cache item
type CacheEntry struct {
	Value      string
	Expiration time.Time
}

// Cache represents an in-memory cache with expiration handling
type Cache struct {
	data map[string]CacheEntry
	mu   sync.Mutex
}

// NewCache creates a new cache
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]CacheEntry),
	}
}

// Set adds a new entry to the cache
func (c *Cache) Set(key string, value string, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry{
		Value:      value,
		Expiration: time.Now().Add(duration),
	}
}

// Get retrieves an entry from the cache
func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, found := c.data[key]
	if !found || entry.Expiration.Before(time.Now()) {
		return "", false
	}
	return entry.Value, true
}

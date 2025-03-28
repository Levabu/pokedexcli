package pokecache

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	data map[string]CacheEntry
	mu sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.data[key] = CacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntry, isFound := c.data[key]
	if !isFound {
		return nil, false
	}
	return cacheEntry.val, true
}

func ConvertCacheEntry[T any](val []byte, data *T) (error) {
	err := json.Unmarshal(val, &data)
	if err != nil {
		return fmt.Errorf("error encoding data: %v", err)
	}
	return nil
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	start := time.Now()
		
	for t := range ticker.C {
		now := time.Now()
		fmt.Printf("%v have passed\n", now.Sub(start))
		start = now
		c.mu.Lock()
		for key, entry := range c.data {
			if t.Sub(entry.createdAt) >= interval {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data: make(map[string]CacheEntry),
		mu: sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}
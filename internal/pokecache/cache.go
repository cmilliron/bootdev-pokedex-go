package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu			sync.Mutex
	cacheEntry	map[string]cacheEntry
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) (*Cache) {
	cache := Cache{
		cacheEntry: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return &cache
} 

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exist := c.cacheEntry[key]
	if exist == false {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cacheEntry {
			if time.Since(entry.createdAt) > interval {
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}

}
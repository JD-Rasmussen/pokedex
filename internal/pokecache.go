package internal

import (
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}
	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			c.reapLoop(interval)
		}
	}()
	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > interval {
			delete(c.entries, key)
		}
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

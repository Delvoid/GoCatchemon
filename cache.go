package main

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mutex   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(reapInterval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mutex:   &sync.Mutex{},
	}
	go cache.reapLoop(reapInterval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, exists := c.entries[key]
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}

}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, entry := range c.entries {
		if now.Sub(entry.createdAt) > interval {
			delete(c.entries, key)
		}
	}
}

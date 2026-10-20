package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mutex    sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries:  map[string]cacheEntry{},
		interval: interval,
	}

	go c.reapLoop()

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	val, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}

	return val.val, true
}

func (c *Cache) reapLoop() {
	reap := func() {
		c.mutex.Lock()
		defer c.mutex.Unlock()

		for k, e := range c.entries {
			if e.createdAt.Add(c.interval).Before(time.Now()) {
				delete(c.entries, k)
			}
		}
	}

	reap()

	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		reap()
	}
}

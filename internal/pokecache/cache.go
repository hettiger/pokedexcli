package pokecache

import (
	"bytes"
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mutex    sync.Mutex
	done     chan struct{}
	stopOnce sync.Once
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

// NewCache creates a cache that removes entries older than interval.
// It panics if interval is not positive.
func NewCache(interval time.Duration) *Cache {
	if interval <= 0 {
		panic("pokecache: interval must be positive")
	}

	c := Cache{
		entries:  map[string]cacheEntry{},
		interval: interval,
		done:     make(chan struct{}),
	}

	go c.reapLoop()

	return &c
}

// Add stores a copy of val under key, replacing any existing entry.
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{
		val:       bytes.Clone(val),
		createdAt: time.Now(),
	}
}

// Get returns a copy of the value stored under key.
// The bool reports whether the key was found.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return bytes.Clone(entry.val), true
}

// Stop terminates the background reaping goroutine. It is safe to call multiple times.
func (c *Cache) Stop() {
	c.stopOnce.Do(func() {
		close(c.done)
	})
}

func (c *Cache) reapLoop() {
	reap := func(now time.Time) {
		c.mutex.Lock()
		defer c.mutex.Unlock()

		for k, e := range c.entries {
			if e.createdAt.Add(c.interval).Before(now) {
				delete(c.entries, k)
			}
		}
	}

	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			reap(time.Now())
		case <-c.done:
			return
		}
	}
}

package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries : make(map[string]cacheEntry),
		interval : interval,
	}
	go cache.reapLoop()
	return cache
}


func (c *Cache) Add(key string, val []byte) {
	
	c.mux.Lock()
	defer c.mux.Unlock()

	c.entries[key] = cacheEntry {
		val: 				val,
		createdAt:  time.Now(),
	}
	return
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	val, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reapLoop() {
	interval := c.interval
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mux.Lock()
		
		timeAgo := time.Now().Add(-interval)

		for key,value := range c.entries {
			if value.createdAt.Before(timeAgo) {
				delete(c.entries, key)
			}
		}
		c.mux.Unlock()
	}
}

package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	mux			 sync.RWMutex
	entries  map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt			time.Time
	val 					[]byte


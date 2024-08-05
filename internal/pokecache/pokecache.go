package pokecache

import (
	"sync"
	"time"
)

const CacheTTL = time.Second * 30

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	items map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(ttl time.Duration) *Cache {
	c := &Cache{
		items: make(map[string]cacheEntry),
	}

	go func() {
		for range time.Tick(ttl) {
			c.mu.Lock()

			for key, item := range c.items {
				if time.Since(item.createdAt) > ttl {
					delete(c.items, key)
				}
			}

			c.mu.Unlock()
		}
	}()

	return c
}

func (c *Cache) Add(key string, apiRes []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = cacheEntry{
		val:       apiRes,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found {
		return item.val, false
	}

	// Not sure if we really need this
	// if item.isExpired() {
	// 	delete(c.items, key)
	// 	return item.val, false
	// }

	return item.val, true
}

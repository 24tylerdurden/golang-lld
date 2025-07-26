package internals

import (
	"sync"
	"time"
)

type Item struct {
	value      interface{}
	expiration int64 // Unix timestamp in seconds
}

func (i Item) isExpired() bool {
	if i.expiration == 0 {
		return false
	}
	return time.Now().Unix() > i.expiration
}

type Cache struct {
	items map[string]Item
	mutex sync.RWMutex
}

func NewCache() *Cache {
	cache := &Cache{
		items: make(map[string]Item),
	}
	go cache.cleanup()

	return cache
}

// Set adds a key-value pair with expiration time (in seconds)

func (c *Cache) Set(key string, value interface{}, expirationSeconds int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var expiration int64

	if expirationSeconds > 0 {
		expiration = time.Now().Unix() + expirationSeconds
	}

	// store the key in the cache
	c.items[key] = Item{
		value:      value,
		expiration: expiration,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	item, exists := c.items[key]
	c.mutex.RUnlock()

	if !exists {
		return nil, false
	}

	if item.isExpired() {
		c.mutex.Lock()
		delete(c.items, key)
		c.mutex.Unlock()
		return nil, false
	}

	return item.value, true
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.items, key)
}

// cleanup periodically removed expired keys
func (c *Cache) cleanup() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			for key, items := range c.items {
				if items.isExpired() {
					delete(c.items, key)
				}
			}
			c.mutex.Unlock()
		case <-time.After(15 * time.Second):
			return
		}
	}

}

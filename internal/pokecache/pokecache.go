package pokecache

import "time"

type Cache struct {
	cache map[string]CacheEntry
}

type CacheEntry struct {
	val       []byte
	createdat time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
	}
	// this cannot be on the main thread
	// this will never exit so nothing else would ever happen
	go c.ReapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = CacheEntry{
		val:       val,
		createdat: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	ce, ok := c.cache[key]
	return ce.val, ok
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Reap(interval)
	}

}
func (c *Cache) Reap(interval time.Duration) {
	minutesAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdat.Before(minutesAgo) {
			delete(c.cache, k)
		}
	}
}

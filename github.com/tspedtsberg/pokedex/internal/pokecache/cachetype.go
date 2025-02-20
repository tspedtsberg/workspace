package pokecache

import (
	"time"
	"sync"
)

//Cache struct
//data ex:
//"https://example.com": {createdAt: 11:59:54, val: []byte("data")},
//"https://another.com": {createdAt: 12:00:00, val: []byte("newer data")
type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

// NewCache
func NewCache(timeout time.Duration) Cache {
	c := Cache{
		cache: 	make(map[string]cacheEntry),
		mux: 	&sync.Mutex{},
	}

	go c.reapLoop(timeout)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}
//Locks and unlock with mutex
//Looks up the cache "c".cache[key] and 
//makes val = the cacheentry struct.
//hence val.val is equal to the []byte data
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

//Creates a ticker. Which ticks on the channel C.
// It ticks every (interval) duration. 
//The for range ticker.C loops over the channel. Every time it ticks it calls the reap function
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}
//Reap func.
//Gets called with the current time and an interval(5 * seconds)
//It locks the cache
//the for loop iterates over the cache, if a cache is create before current time - the interval duration
//it gets deleted
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}


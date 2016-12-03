package updatecache

import (
	"sync"
	"time"
)

type IntCache struct {
	delay       time.Duration
	lock        *sync.Mutex
	updateCache map[int]bool
}

func NewInt(delay time.Duration) *IntCache {
	c := new(IntCache)
	c.delay = delay
	c.updateCache = make(map[int]bool)
	c.lock = new(sync.Mutex)
	return c
}

func (c *IntCache) Add(key int) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	if _, ok := c.updateCache[key]; !ok {
		c.updateCache[key] = false
		go c.remove(key)
		return true
	}

	return false
}

func (c *IntCache) remove(key int) {
	time.Sleep(c.delay)
	c.lock.Lock()
	delete(c.updateCache, key)
	c.lock.Unlock()
}

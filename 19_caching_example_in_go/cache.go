// cache.go
package cache
import "sync" // We’ll use strings as our cache keys
type Key string // And you can store any type of value in our cache
type Value interface{}
type Cache struct 
{    
	data map[Key]Value    
	lock sync.RWMutex
}

const Threaded = true

func (c *Cache) Get (k Key) (Value, bool){
	if Threaded {
		// We're using Read lock/unlock methods
		// since we're not mutating the cache, only
		// reading values from it

		c.lock.RLock()
		defer c.lock.RUnlock()

	}
	value, exists := c.data[k]
	if !exists {
		return nil, false
	}
	return value, true
}

func (c *Cache) Set(k Key, v Value){
	if Threaded {
		// We're using regular Lock/Unlock methods here
		// since we're mutating the cache
		c.lock.Lock()
		defer c.lock.Unlock()
	}
	c.data[k] = v
}

func New() *Cache {
	cache := &Cache{
		data : make(map[Key]Value),
	}
	return cache
}

func (c *Cache) Remove(k Key){
	if Threaded {
		c.lock.Lock()
		defer c.lock.Unlock()
	}

	delete(c.data, k)
}


package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

/*type cacheItem struct {
	key   string
	value interface{}
}*/

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) lastKey() Key {
	var lk Key
	for k, v := range c.items {
		if v == c.queue.Back() {
			lk = k
			break
		}
	}
	return lk
}

func (c *lruCache) Set(key Key, v interface{}) bool {
	if k, exist := c.items[key]; exist {
		k.Value = v
		c.queue.MoveToFront(k)
		return true
	}
	c.items[key] = c.queue.PushFront(v)
	if c.queue.Len() > c.capacity {
		lk := c.lastKey()
		c.queue.Remove(c.items[lk])
		delete(c.items, lk)
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if k, exist := c.items[key]; exist {
		c.queue.MoveToFront(k)
		return c.items[key].Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	for k, v := range c.items {
		c.queue.Remove(v)
		delete(c.items, k)
	}
}

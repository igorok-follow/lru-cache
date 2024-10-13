package cache

import (
	"lru/linked_list"
	"sync"
)

type (
	Cache interface {
		Set(k, v string)
		Get(k string) (string, bool)
	}

	cache struct {
		overflow int
		list     linked_list.List
		hash     map[string]*linked_list.Node
		mu       sync.RWMutex
	}
)

func New(overflow int) Cache {
	return &cache{
		overflow: overflow,
		list:     linked_list.New(),
		hash:     make(map[string]*linked_list.Node),
	}
}

func (c *cache) Set(k, v string) {
	node, ok := c.hash[k]
	if ok {
		node.Val = v

		c.list.Delete(node)
		c.list.Add(node)
	} else {
		node = &linked_list.Node{
			Key: k,
			Val: v,
		}

		c.list.Add(node)

		defer c.mu.Unlock()
		c.mu.Lock()
		c.hash[k] = node

		if len(c.hash) > c.overflow {
			lru := c.list.GetLRU()
			delete(c.hash, lru.Key)
			c.list.Delete(lru)
		}
	}
}

func (c *cache) Get(k string) (string, bool) {
	c.mu.RLock()
	node, ok := c.hash[k]
	if !ok {
		return "", false
	}
	c.mu.RUnlock()

	c.list.Delete(node)
	c.list.Add(node)

	return node.Val, true
}

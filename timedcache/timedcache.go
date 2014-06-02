package timedcache

import (
	"container/list"
	"time"
	"sync"
)

const INIT_CACHE_SIZE = 100 * 1000

type cacheItem struct {
	key        interface{}
	value      interface{}
	expireTime time.Time
}
type TimedCache struct {
 	sync.Mutex
	items          map[interface{}]*list.Element
	expireList     *list.List
	expireInterval time.Duration
	timeToLive     time.Duration
}

func New() *TimedCache {
	c :=  &TimedCache{
		items:          make(map[interface{}]*list.Element, INIT_CACHE_SIZE),
		expireList:     list.New(),
		timeToLive: time.Second*10,
		expireInterval: time.Second,
	}

	go c.expireCheckLoop()

	return c
}

func (c *TimedCache) Set(key interface{}, value interface{}) {
	c.Lock()
	defer c.Unlock()
	var item *cacheItem
	if e, have := c.items[key]; have {
		item, _ = e.Value.(*cacheItem)
		c.expireList.MoveToBack(e)
	} else {
		item = &cacheItem{
			key:key,
			value: value,
		}
		newElem := c.expireList.PushBack(item)
		c.items[key] = newElem
	}

	item.expireTime = time.Now()
}

func (c *TimedCache) Get(key interface{}) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()
	var item *cacheItem
	if e, have := c.items[key]; have {
		item, _ = e.Value.(*cacheItem)
		return item.value, true
	} else {
		return nil, false
	}
}

func (c* TimedCache)expireCheckLoop(){
	t := time.NewTicker(c.expireInterval)
	for{
		select{
		case <-t.C:
			c.expire()
		}
	}
}

func (c *TimedCache)expire(){
	c.Lock()
	defer c.Unlock()
	now := time.Now()
	for e := c.expireList.Front(); e != nil; {
		next := e.Next()
		item := e.Value.(*cacheItem)
		if item.expireTime.After(now){
			break
		}

		c.expireList.Remove(e)
		delete(c.items, item.key)
		e = next
	}
}

package timedcache

import(
    "fmt"
    "time"
    "container/list"
)

const INIT_CACHE_SIZE = 100*1000

type KeyType uint64
type ValueType bool
const DEFAULT_VALUE ValueType = false 

type cacheItem struct{
    value ValueType
    expireTime time.Time
}
type TimedCache struct{
    items map[KeyType] *list.Element
    expireList *list.List
    expireInterval time.Duration
}

func New() *TimedCache{
    return &TimedCache{
        items : make(map[KeyType] *list.Element, INIT_CACHE_SIZE),
        expireList : list.New(),
        expireInterval : time.Second,
    }
}

func (c *TimedCache) Set(key KeyType, value ValueType){
    var item *cacheItem
    if e, have := c.items[key]; have{
        var ok bool
        item, ok = e.Value.(*cacheItem)
        if !ok{
            panic(fmt.Sprintf("Invalid value for timedcache, key=%v", key))
        }
        c.expireList.MoveToBack(e)
    }else{
        item = &cacheItem{
            value : value,
        }
        newElem := c.expireList.PushBack(item)
        c.items[key] = newElem
    }

    item.expireTime = time.Now()
}

func (c *TimedCache) Get(key KeyType) (ValueType, bool){
    var item *cacheItem
    if e, have := c.items[key]; have{
        var ok bool
        item, ok = e.Value.(*cacheItem)
        if !ok{
            panic(fmt.Sprintf("Invalid value for timedcache, key=%v", key))
        }

        return item.value, true
    }else{
        return DEFAULT_VALUE, false
    }
}

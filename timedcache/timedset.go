package timedcache

type TimedSet struct {
	cache *TimedCache
}

func NewSet() *TimedSet {
	return &TimedSet{
		cache: New(),
	}
}

func (ts *TimedSet) Set(key interface{}) {
	ts.cache.Set(key, nil)
}

func (ts *TimedSet) Exist(key interface{}) bool {
	_, exists := ts.cache.Get(key)
	return exists
}

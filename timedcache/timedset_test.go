package timedcache

import (
	"testing"
	"time"
)

func TestSetSet(t *testing.T) {
	cache := NewSet()
	cache.Set(1)
	if !cache.Exist(1) {
		t.Errorf("failed to get value for '1'")
	}
}

func TestSetExpire(t *testing.T){
	cache := NewSet()
	cache.cache.timeToLive = time.Second * 1
	cache.Set(2)
	cache.Set(11)

	if !cache.Exist(2) {
		t.Errorf("failed to get value for '2'")
	}

	time.Sleep(time.Second *2)
	if cache.Exist(2) {
		t.Errorf("failed to get value for '2'")
	}
}


package timedcache

import (
	"testing"
	"time"
	//"fmt"
	//"math/rand"
)

func TestSet(t *testing.T) {
	cache := New()
	cache.Set(1, nil)

	_, ok := cache.Get(1)

	if !ok {
		t.Errorf("failed to get value for '1'")
	}

	time.Sleep(time.Millisecond * 20)

	cache.Get(1) //get does not auto-flash the expiration time

	_, ok = cache.Get(1)

	if !ok {
		t.Errorf("shoud not get value for '1'")
	}

}

func TestExpire(t *testing.T){
	cache := New()
	cache.timeToLive = time.Second * 1
	cache.Set(2, nil)
	cache.Set(11, nil)

	_, ok := cache.Get(2)

	if !ok {
		t.Errorf("failed to get value for '2'")
	}

	time.Sleep(time.Second *2)
	_, ok = cache.Get(2)

	if ok {
		t.Errorf("should not get value for '2'")
	}


}

func BenchmarkAdd(b *testing.B) {
	//long time init
	b.StopTimer()

	cache := New()
	for i := 0; i < 1000*1000; i++ {
		// expireTime :=  time.Millisecond * time.Duration(rand.Intn(1000) + 1)
		cache.Set(i, nil)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		cache.Get(i)
	}
}

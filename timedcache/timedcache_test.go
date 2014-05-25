package timedcache

import  (
    "testing"
    "time"
    //"fmt"
	//"math/rand"
)


func TestPut(t *testing.T){
	cache := New()
	cache.Set(1, true)

	_, ok := cache.Get(1)

	if !ok {
		t.Errorf("failed to get value for '1'")
	}

    time.Sleep(time.Millisecond * 20)

	cache.Get(1) //get does not auto-flash the expiration time

	_, ok =  cache.Get(1)

	if !ok {
		t.Errorf("shoud not get value for '1'")
	}

}

func BenchmarkAdd(b *testing.B){
	//long time init
	b.StopTimer()

	cache := New()
	for i := 0; i < 1000*1000; i++ {
        // expireTime :=  time.Millisecond * time.Duration(rand.Intn(1000) + 1)
		cache.Set(KeyType(i), true)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++{
        cache.Get(KeyType(i))
	}
}

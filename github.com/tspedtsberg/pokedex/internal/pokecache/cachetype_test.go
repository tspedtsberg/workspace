package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://anotherexample.com",
			val: []byte("moretestdata"),
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(v.key, v.val)
			val, ok := cache.Get(v.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(v.val) {
				t.Errorf("expted to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Second
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("exptected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("exptected not to find key")
		return
	}
}


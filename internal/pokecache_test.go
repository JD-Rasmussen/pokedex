package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	const interval = 5 * time.Second
	cache := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cache {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)

			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected key %s to exist in cache", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Expected value %s for key %s, got %s", c.val, c.key, val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("Expected key %s to exist in cache", "https://example.com")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("Expected key %s to be reaped from cache", "https://example.com")
		return
	}
}

func TestGetNonExistentKey(t *testing.T) {
	cache := NewCache(5 * time.Second)
	_, ok := cache.Get("https://nonexistent.com")
	if ok {
		t.Errorf("Expected key %s to not exist in cache", "https://nonexistent.com")
		return
	}
}

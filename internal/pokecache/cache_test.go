package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	c := NewCache(2 * time.Second)

	key := "test-key"
	val := []byte("hello")

	c.Add(key, val)
	got, ok := c.Get(key)

	if !ok {
		t.Fatalf("expected key to be found")
	}

	if string(got) != "hello" {
		t.Fatalf("expected 'hello', got %s", string(got))
	}
}

func TestCacheExpires(t *testing.T) {
	c := NewCache(1 * time.Second)

	key := "old"
	c.Add(key, []byte("data"))

	// wait for entry to expire
	time.Sleep(1500 * time.Millisecond)

	_, ok := c.Get(key)
	if ok {
		t.Fatalf("expected key to expire but it still exists")
	}
}

package pokecache

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

// TestAddGet verifies that we can add an item and retrieve it successfully.
func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)
	cases := []struct {
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
		{
			key : "https://api.example.com/pokemon/1",
			val : []byte("bulbasaur data"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {

			err := cache.Add(c.key, c.val)
			if err != nil {
				t.Fatalf("failed to add item to cache: %v", err)
			}
		
			actual, found := cache.Get(c.key)
			if !found {
				t.Fatalf("expected to find key %q in cache, but it was missing", c.key)
			}
		
			if !bytes.Equal(actual, c.val) {
				t.Errorf("expected value %q, but got %q", string(c.val), string(actual))
			}
		})
	}

}

// TestReapLoop verifies that old cache entries are automatically reaped.
func TestReapLoop(t *testing.T) {
	// Use a tiny interval for fast tests
	const interval = 10 * time.Millisecond
	cache := NewCache(interval)

	key := "expire-me"
	val := []byte("temporary data")

	_ = cache.Add(key, val)

	// Wait longer than the interval duration so the reap loop ticks
	time.Sleep(interval * 2)

	_, found := cache.Get(key)
	if found {
		t.Errorf("expected key %q to be reaped and deleted, but it was still found", key)
	}
}

// TestReapLoopSavesValid ensures items that haven't expired yet remain untouched.
func TestReapLoopSavesValid(t *testing.T) {
	const interval = 50 * time.Millisecond
	cache := NewCache(interval)

	key := "stay-alive"
	val := []byte("fresh data")

	_ = cache.Add(key, val)

	// Wait a tiny bit, but less than the expiration interval
	time.Sleep(10 * time.Millisecond)

	_, found := cache.Get(key)
	if !found {
		t.Errorf("expected key %q to still be valid and alive, but it was prematurely reaped", key)
	}
}

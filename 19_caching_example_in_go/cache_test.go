// cache_test.go

package cache

import (
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	dns := New()
	dns.Set("apple.com", "10.0.0.1")
	ip, exists := dns.Get("apple.com")
	if !exists {
		t.Error("apple.com wasn't found")
	}
	if ip == nil {
		t.Error("dns[apple.com] is nil")
	}
	if ip != "10.0.0.1" {
		t.Error("dns[apple.com] is not same as 10.0.0.1")
	}
}

func TestRemove(t *testing.T) {
	fruits := New()
	fruits.Set("Apple", 1.39)
	applePrice, exists := fruits.Get("Apple")

	if !exists {
		t.Error("Apple price wasn't set")
	}
	if applePrice == nil {
		t.Error("Apple price is nil")
	}
	if applePrice != 1.39 {
		t.Error("Apple price isn't 1.39")
	}

	fruits.Remove("Apple")

	applePrice, exists = fruits.Get("Apple")

	if exists {
		t.Error("Apple price was not removed")
	}

	if applePrice != nil {
		t.Error("Apple price is not nil after removal")
	}
}

func TestThreading(t *testing.T) {
	bloom := New()
	bloom.Set("Gravity", 9.8)

	go func() {
		for {
			_, _ = bloom.Get("Gravity")
		}
	}()

	time.Sleep(1 * time.Second)

	go func() {
		for {
			bloom.Set("Maude", 4.6)
		}
	}()

	time.Sleep(5 * time.Second)
}

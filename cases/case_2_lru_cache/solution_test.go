package case_2_lru_cache

import (
	"testing"
)

func TestBasicGetPut(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")

	if val, ok := cache.Get("a"); !ok || val != "1" {
		t.Errorf("Expected Get('a') = '1', true; got '%s', %v", val, ok)
	}

	if val, ok := cache.Get("b"); !ok || val != "2" {
		t.Errorf("Expected Get('b') = '2', true; got '%s', %v", val, ok)
	}
}

func TestEviction(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")
	cache.Put("c", "3")

	if _, ok := cache.Get("a"); ok {
		t.Errorf("Expected Get('a') to return false; got true")
	}

	if val, ok := cache.Get("b"); !ok || val != "2" {
		t.Errorf("Expected Get('b') = '2', true; got '%s', %v", val, ok)
	}

	if val, ok := cache.Get("c"); !ok || val != "3" {
		t.Errorf("Expected Get('c') = '3', true; got '%s', %v", val, ok)
	}
}

func TestUpdateExistingKey(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	cache.Put("a", "updated")

	if val, ok := cache.Get("a"); !ok || val != "updated" {
		t.Errorf("Expected Get('a') = 'updated', true; got '%s', %v", val, ok)
	}
}

func TestAccessOrderAdjustment(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")

	cache.Get("a")

	cache.Put("c", "3")

	if _, ok := cache.Get("b"); ok {
		t.Errorf("Expected Get('b') to return false; got true")
	}

	if val, ok := cache.Get("a"); !ok || val != "1" {
		t.Errorf("Expected Get('a') = '1', true; got '%s', %v", val, ok)
	}

	if val, ok := cache.Get("c"); !ok || val != "3" {
		t.Errorf("Expected Get('c') = '3', true; got '%s', %v", val, ok)
	}
}

func TestCapacityOne(t *testing.T) {
	cache := NewLRUCache(1)

	cache.Put("a", "1")
	if val, ok := cache.Get("a"); !ok || val != "1" {
		t.Errorf("Expected Get('a') = '1', true; got '%s', %v", val, ok)
	}

	cache.Put("b", "2")
	if _, ok := cache.Get("a"); ok {
		t.Errorf("Expected Get('a') to return false; got true")
	}

	if val, ok := cache.Get("b"); !ok || val != "2" {
		t.Errorf("Expected Get('b') = '2', true; got '%s', %v", val, ok)
	}
}

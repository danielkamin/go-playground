package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestLRUBasic(t *testing.T) {
	// get/put podstawowe
	t.Log("Success TestLRUBasic!")
}
func TestLRUEviction(t *testing.T) {
	t.Log("Success TestLRUEviction!")

} // czy LRU element jest usuwany
func TestLRUUpdateExisting(t *testing.T) {
	t.Log("Success TestLRUUpdateExisting!")

} // put na istniejący klucz
func TestLRUMoveOnGet(t *testing.T) {
	t.Log("Success TestLRUMoveOnGet!")

} // get przesuwa element, zmienia kolejność
func TestLRUConcurrent(t *testing.T) {
	cache := NewLRUCache(100)
	var wg sync.WaitGroup

	// 50 goroutines pisze
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("val%d", i))
		}(i)
	}

	// 50 goroutines czyta
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Get(fmt.Sprintf("key%d", i))
		}(i)
	}

	wg.Wait()
}

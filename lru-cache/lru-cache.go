package main

type LRUCache struct {
	capacity int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
	}
}

func (lc *LRUCache) get(key string) *Node {
	return nil
}
func (lc *LRUCache) put(key, value string) {

}

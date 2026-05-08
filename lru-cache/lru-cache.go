package main

type LRUCache struct {
	capacity  int
	dll       *DLL
	hashTable *HashTable
	len       int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity:  capacity,
		dll:       NewDll(),
		hashTable: NewHashTable(capacity),
		len:       0,
	}
}

func (lc *LRUCache) Get(key string) (string, bool) {
	entry, found := lc.hashTable.get(key)
	if !found {
		return "", false
	}
	lc.dll.remove(entry.value)
	lc.dll.addToFront(entry.value)
	return entry.value.value, true
}
func (lc *LRUCache) Put(key, value string) {
	entry, exists := lc.hashTable.get(key)
	if exists {
		entry.value.value = value
		lc.dll.remove(entry.value)
		lc.dll.addToFront(entry.value)
		return
	}
	newNode := &Node{
		key:   key,
		value: value,
	}
	lc.hashTable.put(key, newNode)
	lc.dll.addToFront(newNode)
	lc.len++

	if lc.len > lc.capacity {
		lastNode := lc.dll.removeLast()
		lc.hashTable.delete(lastNode.key)
		lc.len--
	}
}

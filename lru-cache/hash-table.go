package main

import (
	"hash/fnv"
)

type HashTable struct {
	buckets []*Entry
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*Entry, size),
	}
}

type Entry struct {
	key   string
	value *Node
	next  *Entry
}

func newHashKey(key string, itemsLen int) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int(hash.Sum32()) % itemsLen
}

func (ht *HashTable) get(key string) (*Entry, bool) {
	hashKey := newHashKey(key, len(ht.buckets))
	entry := ht.buckets[hashKey]
	for entry != nil {
		if entry.key == key {
			return entry, true
		}
		entry = entry.next
	}
	return nil, false
}
func (ht *HashTable) put(key string, value *Node) bool {
	hashKey := newHashKey(key, len(ht.buckets))
	if existingEntry, exists := ht.get(key); exists {
		existingEntry.value = value
		return true
	}
	ht.buckets[hashKey] = &Entry{
		key:   key,
		value: value,
		next:  ht.buckets[hashKey],
	}
	return true
}

func (ht *HashTable) delete(key string) bool {
	hashKey := newHashKey(key, len(ht.buckets))
	current := ht.buckets[hashKey]
	var previous *Entry
	for current != nil {
		if current.key == key {
			if previous == nil {
				ht.buckets[hashKey] = current.next
			} else {
				previous.next = current.next
			}
			return true
		}
		previous = current
		current = current.next
	}
	return false
}

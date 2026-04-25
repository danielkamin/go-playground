package main

import "fmt"

func main() {
	ht := NewHashTable(16)
	ht.Put("2026-01-01", &Node{value: "28.9"})
	ht.Put("2026-01-02", &Node{value: "27.9"})
	ht.Put("2026-01-03", &Node{value: "22.9"})
	ht.Put("2026-01-04", &Node{value: "23.9"})
	ht.Put("2026-01-05", &Node{value: "28.9"})
	ht.Put("2026-01-05", &Node{value: "28.9"})
	ht.Put("2026-01-06", &Node{value: "28.9"})

	fmt.Printf("%+v", ht)
}

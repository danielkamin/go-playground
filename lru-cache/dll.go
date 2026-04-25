package main

type DLL struct {
	head *Node
	tail *Node
}

func NewDll() *DLL {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DLL{head, tail}
}

type Node struct {
	key, value string
	prev, next *Node
}

func (dll *DLL) AddToFront(node *Node) {
	node.next = dll.head.next
	node.prev = dll.head
	dll.head.next.prev = node
	dll.head.next = node
}
func (dll *DLL) Remove(node *Node) {

}

func (dll *DLL) RemoveLast() {

}

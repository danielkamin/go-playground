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

func (dll *DLL) addToFront(node *Node) {
	node.next = dll.head.next
	node.prev = dll.head
	dll.head.next.prev = node
	dll.head.next = node
}
func (dll *DLL) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (dll *DLL) removeLast() *Node {
	last := dll.tail.prev
	if last == dll.head {
		return nil
	}
	dll.remove(last)
	return last
}

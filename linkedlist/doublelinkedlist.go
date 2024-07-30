package linkedlist

import "fmt"

// DoubleNode represents a node in the linked list
type DoubleNode[T any] struct {
	Data T
	Prev *DoubleNode[T]
	Next *DoubleNode[T]
}

// DoubleLinkedList represents the linked list
type DoubleLinkedList[T any] struct {
	Head *DoubleNode[T]
	Tail *DoubleNode[T]
}

// Append adds a new node with the given data to the end of the list
func (ll *DoubleLinkedList[T]) Append(data T) *DoubleNode[T] {
	newNode := &DoubleNode[T]{Data: data, Prev: nil, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
		return newNode
	}

	if ll.Tail != nil {
		newNode.Prev = ll.Tail
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}

	return newNode
}

// Push adds a new node with the given data to the head of the list
func (ll *DoubleLinkedList[T]) Push(data T) *DoubleNode[T] {
	newNode := &DoubleNode[T]{Data: data, Prev: nil, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
		return newNode
	}

	newNode.Next = ll.Head
	ll.Head.Prev = newNode
	ll.Head = newNode

	return newNode
}

func (ll *DoubleLinkedList[T]) Remove(node *DoubleNode[T]) *DoubleNode[T] {

	switch {
	case node == ll.Head:
		ll.Head = ll.Head.Next
		ll.Head.Prev = nil
		return node
	case node == ll.Tail:
		ll.Tail = ll.Tail.Prev
		ll.Tail.Next = nil
		return node
	default:
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		return node
	}
}

func (ll *DoubleLinkedList[T]) RemoveTail() *DoubleNode[T] {
	return ll.Remove(ll.Tail)
}

// Print displays all the elements in the linked list
func (ll *DoubleLinkedList[T]) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (ll *DoubleLinkedList[T]) MoveToHead(node *DoubleNode[T]) {
	ll.Remove(node)
	node.Next = ll.Head
	ll.Head.Prev = node
	ll.Head = node
	ll.Head.Prev = nil
}

package linkedlist

import "fmt"

type DoubleLinkedList[T any] interface {
	Head() *DoubleNode[T]
	Tail() *DoubleNode[T]
	// Append inserts new item at tail
	Append(data T) *DoubleNode[T]
	// Push appends new item at head
	Push(data T) *DoubleNode[T]
	Remove(node *DoubleNode[T]) *DoubleNode[T]
	RemoveTail() *DoubleNode[T]
	MoveToHead(node *DoubleNode[T])
	Print()
}

func NewDoubleLinkedList[T any]() DoubleLinkedList[T] {
	list := DoubleLinkedListImpl[T]{}
	return &list
}

// DoubleNode represents a node in the linked list
type DoubleNode[T any] struct {
	Data T
	Prev *DoubleNode[T]
	Next *DoubleNode[T]
}

// DoubleLinkedListImpl represents the linked list
type DoubleLinkedListImpl[T any] struct {
	head *DoubleNode[T]
	tail *DoubleNode[T]
}

func (ll *DoubleLinkedListImpl[T]) Head() *DoubleNode[T] {
	return ll.head
}

func (ll *DoubleLinkedListImpl[T]) Tail() *DoubleNode[T] {
	return ll.tail
}

// Append adds a new node with the given data to the end of the list
func (ll *DoubleLinkedListImpl[T]) Append(data T) *DoubleNode[T] {
	newNode := &DoubleNode[T]{Data: data, Prev: nil, Next: nil}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return newNode
	}

	if ll.tail != nil {
		newNode.Prev = ll.tail
		ll.tail.Next = newNode
		ll.tail = newNode
	}

	return newNode
}

// Push adds a new node with the given data to the head of the list
func (ll *DoubleLinkedListImpl[T]) Push(data T) *DoubleNode[T] {
	newNode := &DoubleNode[T]{Data: data, Prev: nil, Next: nil}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return newNode
	}

	newNode.Next = ll.head
	ll.head.Prev = newNode
	ll.head = newNode

	return newNode
}

func (ll *DoubleLinkedListImpl[T]) Remove(node *DoubleNode[T]) *DoubleNode[T] {

	switch {
	case node == ll.head:
		ll.head = ll.head.Next
		ll.head.Prev = nil
		return node
	case node == ll.tail:
		ll.tail = ll.tail.Prev
		ll.tail.Next = nil
		return node
	default:
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		return node
	}
}

func (ll *DoubleLinkedListImpl[T]) RemoveTail() *DoubleNode[T] {
	return ll.Remove(ll.tail)
}

// Print displays all the elements in the linked list
func (ll *DoubleLinkedListImpl[T]) Print() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (ll *DoubleLinkedListImpl[T]) MoveToHead(node *DoubleNode[T]) {
	ll.Remove(node)
	node.Next = ll.head
	ll.head.Prev = node
	ll.head = node
	ll.head.Prev = nil
}

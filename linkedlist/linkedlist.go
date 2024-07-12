package linkedlist

import "fmt"

// Node represents a node in the linked list
type Node[T any] struct {
	Data T
	Next *Node[T]
}

// LinkedList represents the linked list
type LinkedList[T any] struct {
	Head *Node[T]
}

// Append adds a new node with the given data to the end of the list
func (ll *LinkedList[T]) Append(data T) {
	newNode := &Node[T]{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Print displays all the elements in the linked list
func (ll *LinkedList[T]) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (ll *LinkedList[T]) ReverseLinkedList() {
	var prev *Node[T] = nil
	var ptr *Node[T] = ll.Head
	for ptr != nil {
		var next *Node[T] = ptr.Next
		ptr.Next = prev
		prev = ptr
		if next == nil {
			ll.Head = ptr
		}
		ptr = next
	}
}

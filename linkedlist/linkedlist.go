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
func (ll *LinkedList[T]) Append(data T) *Node[T] {
	newNode := &Node[T]{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return newNode
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return newNode
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

func (ll *LinkedList[T]) Length() int {
	count := 0
	current := ll.Head
	for current != nil {
		current = current.Next
		count++
	}
	return count
}

func (ll *LinkedList[T]) Tail() *Node[T] {
	current := ll.Head
	for current != nil {
		if current.Next == nil {
			return current
		}
		current = current.Next
	}
	return nil
}

func (ll *LinkedList[T]) RemoveNthFromEnd(n int) *Node[T] {
	if n == 0 {
		return nil
	}
	fast := ll.Head // this moves to the end
	slow := ll.Head // this one behind the nth from end, 'n' is 1-based

	for count := 0; count < n; count++ {
		if fast == nil { // list is too short
			return nil
		}
		fast = fast.Next
	}
	if fast == nil { // special case, removing head
		res := ll.Head
		ll.Head = ll.Head.Next
		return res
	}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	res := slow.Next
	slow.Next = slow.Next.Next
	return res
}

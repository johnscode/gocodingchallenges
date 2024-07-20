package linkedlist

import "testing"

func TestLinkedList_Append(t *testing.T) {
	list := LinkedList[int]{}
	data := 10
	list.Append(data)
	if list.Head.Data != data {
		t.Errorf("expected %d, got %d", data, list.Head.Data)
	}
}

func TestReverseLinkedList(t *testing.T) {

	ll := LinkedList[int]{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)

	ll.ReverseLinkedList()

	if ll.Head.Data != 4 {
		t.Errorf("expected %d, got %d", 4, ll.Head.Data)
	}

}

func TestLinkedList_RemoveNthFromEnd(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		ll.Append(4)

		nth := ll.RemoveNthFromEnd(2)

		if nth.Data != 3 {
			t.Errorf("expected %d, got %d", 3, nth.Data)
		}
	})

	t.Run("remove head", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)

		nth := ll.RemoveNthFromEnd(3)

		if nth.Data != 1 {
			t.Errorf("expected %d, got %d", 1, nth.Data)
		}
	})

	t.Run("remove tail", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)

		nth := ll.RemoveNthFromEnd(1)

		if nth.Data != 3 {
			t.Errorf("expected %d, got %d", 3, nth.Data)
		}
	})

	t.Run("bad param", func(t *testing.T) {
		ll := LinkedList[int]{}

		nth := ll.RemoveNthFromEnd(0)

		if nth != nil {
			t.Errorf("expected nil")
		}
	})

	t.Run("too short", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.Append(1)
		ll.Append(2)

		nth := ll.RemoveNthFromEnd(3)

		if nth != nil {
			t.Errorf("expected nil")
		}
	})

	t.Run("empty list", func(t *testing.T) {
		ll := LinkedList[int]{}

		nth := ll.RemoveNthFromEnd(2)

		if nth != nil {
			t.Errorf("expected nil")
		}
	})
}

package linkedlist

import (
	"testing"
)

func createInitialList(data int) *DoubleLinkedList[int] {
	list := DoubleLinkedList[int]{}
	list.Append(data)
	return &list
}
func createTwoItemList(data int, data2 int) *DoubleLinkedList[int] {
	list := createInitialList(data)
	list.Append(data2)
	return list
}
func createThreeItemList(data int, data2 int, data3 int) *DoubleLinkedList[int] {
	list := createTwoItemList(data, data2)
	list.Append(data3)
	return list
}

const data = 10
const data2 = 20
const data3 = 30

func TestDoubleLinkedList_Append(t *testing.T) {

	t.Run("initial insert", func(t *testing.T) {
		list := createInitialList(data)
		if list.Head != list.Tail {
			t.Errorf("expected head = tail")
		}
		if list.Head.Prev != nil {
			t.Errorf("expected head.prev to be nil")
		}
		if list.Tail.Next != nil {
			t.Errorf("expected tail.next to be nil")
		}
		if list.Head.Data != data {
			t.Errorf("expected %d, got %d", data, list.Head.Data)
		}
	})
	t.Run("second insert", func(t *testing.T) {
		list := createTwoItemList(data, data2)
		if list.Head.Next != list.Tail {
			t.Errorf("expected head.next to be tail")
		}
		if list.Tail.Prev != list.Head {
			t.Errorf("expected tail.prev to be head")
		}
		if list.Tail.Data != data2 {
			t.Errorf("expected %d, got %d", data2, list.Tail.Data)
		}
	})
	t.Run("third insert", func(t *testing.T) {
		list := createThreeItemList(data, data2, data3)
		if list.Head.Next == list.Tail {
			t.Errorf("expected head.next != tail")
		}
		if list.Tail.Prev == list.Head {
			t.Errorf("expected tail.prev != head")
		}
		if list.Tail.Data != data3 {
			t.Errorf("expected %d, got %d", data3, list.Tail.Data)
		}
	})
}

func TestDoubleLinkedList_Remove(t *testing.T) {
	t.Run("remove head", func(t *testing.T) {
		list := createThreeItemList(data, data2, data3)
		origHead := list.Head
		origTail := list.Tail
		newHead := list.Head.Next
		nodeRemoved := list.Remove(origHead)
		if nodeRemoved != origHead {
			t.Errorf("expected original head returned")
		}
		if list.Head != newHead {
			t.Errorf("expected new head")
		}
		if list.Tail != origTail {
			t.Errorf("expected tail unchanged")
		}
		if list.Head.Next != list.Tail {
			t.Errorf("expected head.next == tail")
		}
		if list.Tail.Prev != list.Head {
			t.Errorf("expected tail.prev == head")
		}
		if nodeRemoved.Data != data {
			t.Errorf("expected %d, got %d", data, nodeRemoved.Data)
		}
	})
	t.Run("remove tail", func(t *testing.T) {
		list := createThreeItemList(data, data2, data3)
		origHead := list.Head
		origTail := list.Tail
		newTail := list.Tail.Prev
		nodeRemoved := list.Remove(origTail)
		if nodeRemoved != origTail {
			t.Errorf("expected original tail returned")
		}
		if list.Tail != newTail {
			t.Errorf("expected new tail")
		}
		if list.Head != origHead {
			t.Errorf("expected head unchanged")
		}
		if list.Head.Next != list.Tail {
			t.Errorf("expected head.next == tail")
		}
		if list.Tail.Prev != list.Head {
			t.Errorf("expected tail.prev == head")
		}
		if nodeRemoved.Data != data3 {
			t.Errorf("expected %d, got %d", data3, nodeRemoved.Data)
		}
	})
	t.Run("remove mid", func(t *testing.T) {
		list := createThreeItemList(data, data2, data3)
		origHead := list.Head
		origTail := list.Tail
		expectedRemoved := list.Head.Next
		nodeRemoved := list.Remove(expectedRemoved)
		if nodeRemoved != expectedRemoved {
			t.Errorf("expected original mid node returned")
		}
		if list.Tail != origTail {
			t.Errorf("expected tail unchanged")
		}
		if list.Head != origHead {
			t.Errorf("expected head unchanged")
		}
		if list.Head.Next != list.Tail {
			t.Errorf("expected head.next == tail")
		}
		if list.Tail.Prev != list.Head {
			t.Errorf("expected tail.prev == head")
		}
		if nodeRemoved.Data != data2 {
			t.Errorf("expected %d, got %d", data2, nodeRemoved.Data)
		}
	})
}

func TestDoubleLinkedList_Push(t *testing.T) {
	list := createInitialList(data)
	origHead := list.Head
	newHead := list.Push(data2)
	if list.Head != newHead {
		t.Errorf("expected new head")
	}
	if list.Head.Next != origHead {
		t.Errorf("expected old head to be second")
	}
	if origHead.Prev != list.Head {
		t.Errorf("expected old head.prev to be new head")
	}
	if list.Head.Data != data2 {
		t.Errorf("expected %d, got %d", data2, list.Head.Data)
	}
}

func TestDoubleLinkedList_RemoveTail(t *testing.T) {
	list := createThreeItemList(data, data2, data3)
	origHead := list.Head
	origTail := list.Tail
	newTail := list.Tail.Prev
	nodeRemoved := list.RemoveTail()
	if nodeRemoved != origTail {
		t.Errorf("expected original tail returned")
	}
	if list.Tail != newTail {
		t.Errorf("expected new tail")
	}
	if list.Head != origHead {
		t.Errorf("expected head unchanged")
	}
	if list.Head.Next != list.Tail {
		t.Errorf("expected head.next == tail")
	}
	if list.Tail.Prev != list.Head {
		t.Errorf("expected tail.prev == head")
	}
	if nodeRemoved.Data != data3 {
		t.Errorf("expected %d, got %d", data3, nodeRemoved.Data)
	}
}

func TestDoubleLinkedList_MoveToHead(t *testing.T) {
	list := createThreeItemList(data, data2, data3)
	origHead := list.Head
	origTail := list.Tail
	nodeToMove := list.Head.Next
	list.MoveToHead(nodeToMove)
	if list.Head != nodeToMove {
		t.Errorf("expected new head")
	}
	if list.Tail != origTail {
		t.Errorf("expected same tail")
	}
	if list.Head.Next != origHead {
		t.Errorf("expected head.next == tail")
	}
	if list.Tail.Prev != list.Head {
		t.Errorf("expected tail.prev == head")
	}
	if list.Head.Data != data2 {
		t.Errorf("expected %d, got %d", data2, list.Head.Data)
	}
}

func TestDoubleLinkedList_Print(t *testing.T) {
	list := createThreeItemList(data, data2, data3)
	list.Print()
}

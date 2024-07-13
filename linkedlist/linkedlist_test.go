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

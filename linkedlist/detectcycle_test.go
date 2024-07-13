package linkedlist

import "testing"

func TestDetectCycle(t *testing.T) {

	t.Run("expect cycle", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.Append(1)
		second := ll.Append(2)
		ll.Append(3)
		ll.Append(4)
		tail := ll.Append(5)

		tail.Next = second // create cycle

		result := DetectCycle(&ll)

		if !result {
			t.Errorf("expected cycle")
		}
	})

	t.Run("expect cycle 2", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		ll.Append(4)
		tail := ll.Append(5)

		tail.Next = ll.Head // create cycle

		result := DetectCycle(&ll)

		if !result {
			t.Errorf("expected cycle")
		}
	})

	t.Run("no cycle", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		ll.Append(4)
		ll.Append(5)

		result := DetectCycle(&ll)

		if result {
			t.Errorf("unexpected cycle")
		}
	})

	t.Run("no cycle on empty", func(t *testing.T) {
		ll := LinkedList[int]{}

		result := DetectCycle(&ll)

		if result {
			t.Errorf("unexpected cycle")
		}
	})

}

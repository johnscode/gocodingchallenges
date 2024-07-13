package linkedlist

import "testing"

func TestMergeLists(t *testing.T) {

	t.Run("actual merge of 2 lists", func(t *testing.T) {
		ll1 := LinkedList[int]{}
		ll1.Append(1)
		ll1.Append(2)
		ll1.Append(4)

		ll2 := LinkedList[int]{}
		ll2.Append(1)
		ll2.Append(3)
		ll2.Append(4)

		result := mergeSortedLists(ll1, ll2)
		length := result.Length()
		if length != 6 {
			t.Errorf("expected length of %d, got %d", 6, length)
		}
	})

	t.Run("2 empty lists", func(t *testing.T) {
		ll1 := LinkedList[int]{}
		ll2 := LinkedList[int]{}
		result := mergeSortedLists(ll1, ll2)
		length := result.Length()
		if length != 0 {
			t.Errorf("expected length of %d, got %d", 0, length)
		}
	})

	t.Run("1 empty list", func(t *testing.T) {
		ll1 := LinkedList[int]{}
		ll2 := LinkedList[int]{}
		ll2.Append(0)
		result := mergeSortedLists(ll1, ll2)
		length := result.Length()
		if length != 1 {
			t.Errorf("expected length of %d, got %d", 1, length)
		}
	})

}

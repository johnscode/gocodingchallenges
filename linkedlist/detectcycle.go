package linkedlist

func DetectCycle[T any](ll *LinkedList[T]) bool {
	slow := ll.Head
	fast := ll.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

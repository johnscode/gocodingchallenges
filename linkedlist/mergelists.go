package linkedlist

func mergeSortedLists(ll1 LinkedList[int], ll2 LinkedList[int]) LinkedList[int] {
	result := LinkedList[int]{}

	p1 := ll1.Head
	p2 := ll2.Head
	rp := &Node[int]{} // dummy node as result head
	result.Head = rp
	for p1 != nil && p2 != nil {
		if p1.Data >= p2.Data {
			rp.Next = p2
			p2 = p2.Next
		} else {
			rp.Next = p1
			p1 = p1.Next
		}
		rp = rp.Next
	}
	if p1 != nil {
		rp.Next = p1
	}
	if p2 != nil {
		rp.Next = p2
	}
	result.Head = result.Head.Next
	return result
}

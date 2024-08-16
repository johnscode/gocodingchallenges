package queue

import "testing"

func makeThreeElementQ() Queue[int] {
	q := MakeQueue[int](10)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	return q
}

func TestQueue(t *testing.T) {
	t.Run("add 3, pop1", func(t *testing.T) {
		q := makeThreeElementQ()
		data := q.Pop()
		if data != 1 {
			t.Errorf("expected %d, got %d", 1, data)
		}
	})
	t.Run("add 3, pop3", func(t *testing.T) {
		q := makeThreeElementQ()
		data := q.Pop()
		data = q.Pop()
		data = q.Pop()
		if data != 3 {
			t.Errorf("expected %d, got %d", 3, data)
		}
	})
}

func TestQueueImpl_Length(t *testing.T) {
	q := makeThreeElementQ()
	length := q.Length()
	if length != 3 {
		t.Errorf("expected %d, got %d", 3, length)
	}
}

package queue

import (
	"gocoding/linkedlist"
)

type Queue[T any] interface {
	Push(data T)
	Pop() T
	Capacity() int
	Length() int
}

type QueueImpl[T any] struct {
	queue    linkedlist.DoubleLinkedList[T]
	capacity int
}

func MakeQueue[T any](capacity int) Queue[T] {
	q := QueueImpl[T]{
		capacity: capacity,
		queue:    linkedlist.NewDoubleLinkedList[T](),
	}
	return &q
}

func (q *QueueImpl[T]) Push(data T) {
	q.queue.Push(data)
}

func (q QueueImpl[T]) Pop() T {
	node := q.queue.RemoveTail()
	return node.Data
}

func (q QueueImpl[T]) Capacity() int {
	return q.capacity
}

func (q QueueImpl[T]) Length() int {
	return q.queue.Length()
}

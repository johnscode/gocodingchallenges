package lrucache

import (
	"gocoding/linkedlist"
	"sync"
)

type LRUCache[T any] interface {
	Get(key string) (value T, found bool)

	Put(key string, value T)

	Keys() []string

	Remove(key string) bool

	Clear()

	Capacity() int

	Len() int
}

type lruCacheEntry[T any] struct {
	key   string
	value T
}

type lruCache[T any] struct {
	capacity int
	keyMap   map[string]*linkedlist.DoubleNode[lruCacheEntry[T]]
	list     linkedlist.DoubleLinkedList[lruCacheEntry[T]]
	mutex    sync.RWMutex
}

func NewLRUCache[T any](capacity int) LRUCache[T] {
	lru := lruCache[T]{capacity: capacity}
	lru.keyMap = map[string]*linkedlist.DoubleNode[lruCacheEntry[T]]{}
	lru.list = linkedlist.NewDoubleLinkedList[lruCacheEntry[T]]()
	return &lru
}

func (l *lruCache[T]) Get(key string) (value T, found bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if node, ok := l.keyMap[key]; ok {
		l.list.MoveToHead(node)
		return node.Data.value, ok
	}
	var zero T
	return zero, false
}

func (l *lruCache[T]) Put(key string, value T) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if node, ok := l.keyMap[key]; ok {
		node.Data = lruCacheEntry[T]{
			key:   key,
			value: value,
		}
		// move the element to the most recent position
		l.list.MoveToHead(node)
	} else {
		// insert the new element at the head
		newNode := l.list.Push(lruCacheEntry[T]{
			key:   key,
			value: value,
		})
		l.keyMap[key] = newNode
	}
	// is eviction necessary
	if len(l.keyMap) > l.capacity {
		nodeRemoved := l.list.RemoveTail()
		delete(l.keyMap, nodeRemoved.Data.key)
	}
}

func (l *lruCache[T]) Keys() []string {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	keys := make([]string, 0, len(l.keyMap))
	current := l.list.Head()
	for current != nil {
		keys = append(keys, current.Data.key)
		current = current.Next
	}
	return keys
}

func (l *lruCache[T]) Remove(key string) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if nodeToRemove, ok := l.keyMap[key]; ok {
		l.list.Remove(nodeToRemove)
		delete(l.keyMap, key)
		return true
	}
	return false
}

func (l *lruCache[T]) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.keyMap = map[string]*linkedlist.DoubleNode[lruCacheEntry[T]]{}
	l.list = &linkedlist.DoubleLinkedListImpl[lruCacheEntry[T]]{}
}

func (l *lruCache[T]) Capacity() int {
	return l.capacity
}

func (l *lruCache[T]) Len() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return len(l.keyMap)
}

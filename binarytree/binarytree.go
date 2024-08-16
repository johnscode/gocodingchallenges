package binarytree

import (
	"gocoding/queue"
	"golang.org/x/exp/constraints"
)

func MakeBinaryTree[T constraints.Ordered]() BinaryTree[T] {
	return &BinaryTreeImpl[T]{}
}

type BinaryTree[T any] interface {
	Insert(data T)
	LevelSearch(data T)
	InorderSearch(data T)
	PreorderSearch(data T)
	PostorderSearch(data T)
}

type BTreeNode[T any] struct {
	Data  T
	Left  *BTreeNode[T]
	Right *BTreeNode[T]
}

type BinaryTreeImpl[T constraints.Ordered] struct {
	Root *BTreeNode[T]
}

func (b *BinaryTreeImpl[T]) Insert(data T) {
	if b.Root == nil {
		b.Root = &BTreeNode[T]{Data: data, Left: nil, Right: nil}
		return
	}
	b.insertRecursive(b.Root, data)
}

func (b *BinaryTreeImpl[T]) insertRecursive(node *BTreeNode[T], data T) {
	if data < node.Data {
		if node.Left == nil {
			node.Left = &BTreeNode[T]{Data: data, Left: nil, Right: nil}
		} else {
			b.insertRecursive(node.Left, data)
		}
	} else if data > node.Data {
		if node.Right == nil {
			node.Right = &BTreeNode[T]{Data: data, Left: nil, Right: nil}
		} else {
			b.insertRecursive(node.Right, data)
		}
	}
}

func (b *BinaryTreeImpl[T]) LevelSearch(data T) {
	q := queue.MakeQueue[*BTreeNode[T]](100)
	q.Push(b.Root)
	for q.Length() > 0 {
		current := q.Pop()
		if current.Left != nil {
			q.Push(current.Left)
		}
		if current.Right != nil {
			q.Push(current.Right)
		}
		println("node data", current.Data)
	}
}

func (b *BinaryTreeImpl[T]) InorderSearch(data T) {
	b.inorderRecurse(b.Root, data)
}

func (b *BinaryTreeImpl[T]) inorderRecurse(node *BTreeNode[T], data T) {
	if node != nil {
		b.inorderRecurse(node.Left, data)
		println("node ", node.Data)
		b.inorderRecurse(node.Right, data)
	}
}

func (b *BinaryTreeImpl[T]) PreorderSearch(data T) {
	b.preorderRecurse(b.Root, data)
}

func (b *BinaryTreeImpl[T]) preorderRecurse(node *BTreeNode[T], data T) {
	if node != nil {
		b.preorderRecurse(node.Left, data)
		println("node ", node.Data)
		b.preorderRecurse(node.Right, data)
	}
}

func (b *BinaryTreeImpl[T]) PostorderSearch(data T) {
	b.postorderRecurse(b.Root, data)
}

func (b *BinaryTreeImpl[T]) postorderRecurse(node *BTreeNode[T], data T) {
	if node != nil {
		b.postorderRecurse(node.Left, data)
		b.postorderRecurse(node.Right, data)
		println("node ", node.Data)
	}
}

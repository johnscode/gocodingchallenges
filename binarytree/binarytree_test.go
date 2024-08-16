package binarytree

import "testing"

func makeBinaryTree() BinaryTree[int] {
	tree := MakeBinaryTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(8)
	return tree
}

func TestBinaryTree(t *testing.T) {
	tree := makeBinaryTree()
	tree.LevelSearch(1)
}

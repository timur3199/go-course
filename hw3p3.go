package main

import (
	"fmt"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type BST struct {
	root *TreeNode
}

func NewBST() *BST {
	return &BST{root: nil}
}

func (bst *BST) Insert(value int) {
	bst.root = bst.insert(bst.root, value)
}

func (bst *BST) insert(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{value: value}
	}

	if value < node.value {
		node.left = bst.insert(node.left, value)
	} else if value > node.value {
		node.right = bst.insert(node.right, value)
	}
	return node
}

func (bst *BST) Remove(value int) {
	bst.root = bst.remove(bst.root, value)
}

func (bst *BST) remove(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}

	if value < node.value {
		node.left = bst.remove(node.left, value)
	} else if value > node.value {
		node.right = bst.remove(node.right, value)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		minNode := bst.findMin(node.right)
		node.value = minNode.value
		node.right = bst.remove(node.right, minNode.value)
	}
	return node
}

func (bst *BST) Find(value int) bool {
	return bst.find(bst.root, value)
}

func (bst *BST) find(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}

	if value == node.value {
		return true
	} else if value < node.value {
		return bst.find(node.left, value)
	} else {
		return bst.find(node.right, value)
	}
}

func (bst *BST) Depth() int {
	return bst.depth(bst.root)
}

func (bst *BST) depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftDepth := bst.depth(node.left)
	rightDepth := bst.depth(node.right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func (bst *BST) findMin(node *TreeNode) *TreeNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

func main() {
	bst := NewBST()
	
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	
	fmt.Printf("Глубина: %d\n", bst.Depth())
	fmt.Printf("Найден 30: %t\n", bst.Find(30))
	fmt.Printf("Найден 100: %t\n", bst.Find(100))
	
	bst.Remove(30)
	fmt.Printf("Найден 30 после удаления: %t\n", bst.Find(30))
	fmt.Printf("Глубина после удаления: %d\n", bst.Depth())
}
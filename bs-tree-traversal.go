package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Parent *TreeNode
	Left *TreeNode
	Right *TreeNode
}

func (n *TreeNode) predecessor(i int) *int {
	node := n.findNode(i)
	if node == nil {
		return nil
	}
	if node.Left != nil {
		return node.Left.max()
	}
	parent := node.Parent
	for parent != nil && node == parent.Left {
		// search for the parent that is less than the current node
		// that is while the node is the left child, its less than its parent
		node = parent
		parent = parent.Parent
	}
	if parent == nil {
		return nil
	}
	return &parent.Val
}

func (n *TreeNode) successor(i int) *int {
	node := n.findNode(i)
	if node == nil {
		return nil
	}
	if node.Right != nil {
		return node.Right.min()
	}
	parent := node.Parent
	for parent != nil && node == parent.Right { 
		//find the first parent that is greater than the node
		//that is, while the node is the right child of its parent, the node is greater than it
		node = parent
		parent = parent.Parent
	}
	if parent == nil {
		return nil
	}
	return &parent.Val
}

func (n *TreeNode) sorted() []int {
	a := make([]int, 0)
	n.appendToArray(&a)
	return a
}

func (n *TreeNode) appendToArray(a *[]int) {
	if n == nil {
		return
	}
	n.Left.appendToArray(a)
	*a = append(*a, n.Val)
	n.Right.appendToArray(a)
}

func (n *TreeNode) findNode(i int) *TreeNode {
	if n == nil {
		return nil
	}
	if n.Val < i {
		return n.Right.findNode(i)
	} else if n.Val > i {
		return n.Left.findNode(i)
	}
	return n
}

func (n *TreeNode) min() *int {
	var val *int
	node := n
	for node != nil {
		val = &node.Val
		node = node.Left
	}
	return val
}

func (n *TreeNode) max() *int {
	var val *int
	node := n
	for node != nil {
		val = &node.Val
		node = node.Right
	}
	return val
}

func (n *TreeNode) find(i int) bool {
	if n == nil {
		return false
	}
	
	if n.Val < i {
		return n.Right.find(i)
	} else if n.Val > i {
		return n.Left.find(i)
	}
	return true
}

func (n *TreeNode) String() string {
	var parent, left, right int
	if n.Parent != nil {
		parent = n.Parent.Val
	}
	if n.Left != nil {
		left = n.Left.Val
	}
	if n.Right != nil {
		right = n.Right.Val
	}
	return fmt.Sprintf("Val: %d Parent: %d Left: %d Right: %d", n.Val, parent, left, right)
}

func (n *TreeNode) traverseInOrder() {
	if n == nil {
		return
	}
	n.Left.traverseInOrder()
	fmt.Println(n)
	n.Right.traverseInOrder()
}

func (n *TreeNode) traversePreOrder() {
	if n == nil {
		return	
	}
	fmt.Println(n)
	n.Left.traversePreOrder()
	n.Right.traversePreOrder()
}

func (n *TreeNode) traversePostOrder() {
	if n == nil {
		return	
	}
	n.Left.traversePreOrder()
	n.Right.traversePreOrder()
	fmt.Println(n)
}

func NewTree(a []int) *TreeNode {
	var root *TreeNode
	for _, i := range a {
		root = addNode(root, nil, i)
	}
	return root
}

func addNode(node *TreeNode, parent *TreeNode, i int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: i, Parent: parent}
	}
	if node.Val < i {
		node.Right = addNode(node.Right, node, i)
	} else {
		node.Left = addNode(node.Left, node, i)	
	}
	return node
}

func run() {
	fmt.Println("go!")
	root := NewTree([]int{4, 6, 7, 2, 6, 8, 9, 2})

	fmt.Println("In-Order")
	root.traverseInOrder()

	fmt.Println("Pre-Order")
	root.traversePreOrder()

	fmt.Println("Post-Order")
	root.traversePostOrder()

	exist := root.find(2)
	fmt.Println(exist)

	exist = root.find(3)
	fmt.Println(exist)

	fmt.Println(*root.min())
	fmt.Println(*root.max())

	root = NewTree([]int{4, 6, 7, 2, 8, 9, 10, 1})
	sorted := root.sorted()
	fmt.Println("sorted", sorted)
	for i := 0; i < len(sorted)-1; i++ {
		fmt.Printf("successor of %d: %d\n", sorted[i], *root.successor(sorted[i]))
	}
	fmt.Printf("successor of %d: %v\n", sorted[len(sorted)-1], root.successor(sorted[len(sorted)-1]))

	for i := len(sorted)-1; i > 0; i-- {
		fmt.Printf("predecessor of %d: %d\n", sorted[i], *root.predecessor(sorted[i]))
	}
	//fmt.Printf("predecessor of %d: %v\n", sorted[0], root.successor(sorted[0]))
}

func main() {
	run()
}

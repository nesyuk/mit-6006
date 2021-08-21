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

func (n *TreeNode) print() {
	if n == nil {
		return
	}
	n.Left.print()
	fmt.Println(n.Val)
	n.Right.print()
}

func (n *TreeNode) String() string {
	return fmt.Sprintf("%d", n.Val)
}

func deleteNode(root *TreeNode, v int) *TreeNode {
	node := findNode(root, v)
	if node == nil {
		return root
	}
	if node.Left == nil {
		root = transplant(root, node, node.Right)
	} else if node.Right == nil {
		root = transplant(root, node, node.Left)	
	} else {
		successor := findMin(node.Right)
		if successor.Parent != node {
			root = transplant(root, successor, successor.Right)
			successor.Right = node.Right
			successor.Right.Parent = successor
		}
		root = transplant(root, node, successor)
		successor.Left = node.Left
		successor.Left.Parent = successor
	}

	return root
}

func transplant(root *TreeNode, node *TreeNode, newNode *TreeNode) *TreeNode {
	if node.Parent == nil {
		root = newNode
	} else if node.Parent.Left == node {
		node.Parent.Left = newNode
	} else {
		node.Parent.Right = newNode
	}
	if newNode != nil {
		newNode.Parent = node.Parent
	}
	return root
}

func findMin(node *TreeNode) *TreeNode {
	n := node
	for n.Left != nil {
		n = n.Left	
	}
	return n
}

func findNode(node *TreeNode, v int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val < v {
		return findNode(node.Right, v)
		
	} else if node.Val > v {
		return findNode(node.Left, v)
	}
	return node
}

func addNode(n *TreeNode, parent *TreeNode, v int) *TreeNode {
	if n == nil {
		return &TreeNode{Val: v, Parent: parent}
	}
	if n.Val < v {
		n.Right = addNode(n.Right, n, v)
	} else {
		n.Left = addNode(n.Left, n, v)
	}
	return n
}

func NewTree(a []int) *TreeNode {
	var root *TreeNode
	for _, v := range a {
		root = addNode(root, nil, v)
	}
	return root
}

func run() {
	fmt.Println("go!")
	r := NewTree([]int{5, 6, 2, 7, 8})
	for _, n := range []int{7, 8, 5, 2, 6} {
		r = deleteNode(r, n)
		fmt.Printf("after deletion of %d\n", n)
		r.print()
	}
}

func main() {
	run()
}

package main

import (
	"fmt"
)

type BSTree struct {
	root *Node
	length int
}

func (t *BSTree) Build(a []int) {
	for _, v := range a {
		t.root = t.add(t.root, v)
		t.length = t.length + 1
	}
}

func (t *BSTree) GetMin() *int {
	node := t.root
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return &node.v
}

func (t *BSTree) Print() {
	t.print(t.root)
}

func (t *BSTree) print(n *Node) {
	if n != nil {
		t.print(n.left)
		fmt.Println(n.v)
		t.print(n.right)
	} else {
		return
	}
}

func (t *BSTree) add(n *Node, v int) *Node {
	if n == nil {
		n = &Node{v: v}
	} else if n.v < v {
		n.right = t.add(n.right, v)
	} else {
		n.left = t.add(n.left, v)
	}
	return n
}

type Node struct {
	left *Node
	right *Node
	v int
}

func (n Node) String() string {
	return fmt.Sprintf("value: %d, left: %v, right: %v", n.v, n.left, n.right)	
}

func run() {
	fmt.Println("go!")

	bstree := BSTree{}
	bstree.Build([]int{1, 2, 3, 4, 5, 6})
	bstree.Print()
	fmt.Printf("%v\n", *bstree.GetMin())
}

func main() {
	run()
}

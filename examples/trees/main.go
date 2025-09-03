package main

import "fmt"

type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(key int) *TreeNode {
	return &TreeNode{Key: key}
}

func main() {
	node0 := NewTreeNode(3)
	node1 := NewTreeNode(4)
	node2 := NewTreeNode(5)

	node0.Left = node1
	node0.Right = node2

	fmt.Println(node0.Key)
	fmt.Println(node0.Left.Key)
	fmt.Println(node0.Right.Key)
}

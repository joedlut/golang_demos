package main

import (
	"fmt"
)

type TreeNode struct {
	data      string
	leftNode  *TreeNode
	rightNode *TreeNode
}

//先序遍历

func (tree *TreeNode) PreOrder() {
	//递归退出的条件
	if tree == nil {
		return
	}
	fmt.Printf("%s ", tree.data)
	tree.leftNode.PreOrder()
	tree.rightNode.PreOrder()
}

//后序遍历

func (tree *TreeNode) PostOrder() {
	if tree == nil {
		return
	}
	tree.leftNode.PostOrder()
	tree.rightNode.PostOrder()
	fmt.Printf("%s ", tree.data)
}

//中序遍历

func (tree *TreeNode) MidOrder() {
	if tree == nil {
		return
	}
	tree.leftNode.MidOrder()
	fmt.Printf("%s ", tree.data)
	tree.rightNode.MidOrder()
}

func main() {
	node1 := TreeNode{
		data:      "1",
		leftNode:  nil,
		rightNode: nil,
	}

	node2 := TreeNode{
		data:      "2",
		leftNode:  nil,
		rightNode: nil,
	}

	node3 := TreeNode{
		data:      "3",
		leftNode:  &node1,
		rightNode: &node2,
	}

	node5 := TreeNode{
		data:      "5",
		leftNode:  nil,
		rightNode: nil,
	}

	node4 := TreeNode{
		data:      "4",
		leftNode:  &node3,
		rightNode: &node5,
	}
	/*
	       4
	      /\
	     3  5
	    / \
	   1   2
	*/

	node4.PreOrder()
	fmt.Println()
	node4.MidOrder()
	fmt.Println()
	node4.PostOrder()

}

package main

import "fmt"

type Comparable func(c1 interface{}, c2 interface{}) bool

type BinaryTree struct {
	node    interface{}
	left    *BinaryTree
	right   *BinaryTree
	lessFun Comparable
}

func New(cmp Comparable) *BinaryTree {
	return &BinaryTree{
		node:    nil,
		lessFun: cmp,
	}
}

func (tree *BinaryTree) Insert(value interface{}) {
	if tree.node == nil {
		tree.node = value
		tree.right = New(tree.lessFun)
		tree.left = New(tree.lessFun)
		return
	}
	if tree.lessFun(value, tree.node) {
		tree.left.Insert(value)
	} else {
		tree.right.Insert(value)
	}
}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
	if tree.node == nil {
		return nil
	}
	if tree.node == value {
		return tree
	}
	if tree.lessFun(value, tree.node) {
		return tree.left.Search(value)
	} else {
		return tree.right.Search(value)
	}
}

func (tree *BinaryTree) Max() interface{} {
	if tree.node == nil || tree.right.node == nil {
		return tree.node
	}
	//return tree.Max()
	return tree.right.Max()
}

func (tree *BinaryTree) Min() interface{} {
	if tree.node == nil || tree.left.node == nil {
		return tree.node
	}
	//return tree.Min()
	return tree.left.Min()
}

func main() {
	var btr *BinaryTree = New(func(c1 interface{}, c2 interface{}) bool {
		//断言
		valueC1, _ := c1.(int)
		valueC2, _ := c2.(int)
		return valueC1 < valueC2
	})

	/*
		12
		/ \
		1 100
		  /\
		23  230
	*/
	btr.Insert(12)
	btr.Insert(100)
	btr.Insert(23)
	btr.Insert(230)
	btr.Insert(1)

	//fmt.Println(btr.right.right.right)

	fmt.Println(btr.Max())

	fmt.Println(btr.Min())

	//fmt.Println(btr.Search(100).node)

	fmt.Println(btr.Search(100))

}

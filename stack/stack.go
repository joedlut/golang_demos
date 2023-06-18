package main

import "fmt"

//LIFO  stack 先进后出

type StackNode struct {
	data interface{}
	next *StackNode
}

type Stack struct {
	sp    *StackNode
	depth int
}

func New() *Stack {
	return &Stack{
		sp:    nil,
		depth: 0,
	}
}

func (stack *Stack) Push(item interface{}) {
	stack.sp = &StackNode{data: item, next: stack.sp}
	stack.depth++
}

func (stack *Stack) Pop() interface{} {
	if stack.depth > 0 {
		data := stack.sp.data
		stack.sp = stack.sp.next
		stack.depth--
		return data
	}
	return nil
}

func (stack *Stack) Peek() interface{} {
	if stack.depth > 0 {
		return stack.sp.data
	}
	return nil
}

func main() {
	stack := New()
	stack.Push(1)
	stack.Push(3)
	stack.Push(5)

	fmt.Println(stack.Peek())

	depth := stack.depth
	for i := 1; i <= depth; i++ {
		fmt.Println(stack.Pop())
	}

	fmt.Println(stack.depth, stack.sp)

}

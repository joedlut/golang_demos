package main

import (
	"fmt"
	"math/rand"
)

type Heap struct {
	data []int
	//堆中可以存储的最大个数
	maxNum int
	//堆中现在数据的个数
	count int
}

func New(capacity int) *Heap {
	return &Heap{
		count:  0,
		maxNum: capacity,
		data:   make([]int, capacity),
	}
}

func (heap *Heap) Insert(n int) {
	if heap.count >= heap.maxNum {
		panic("堆满")
	}
	//下标为0的结点不存储
	heap.count++
	heap.data[heap.count] = n
	i := heap.count
	//i 有父结点 ; 大顶堆
	for i/2 > 0 && heap.data[i] > heap.data[i/2] {
		heap.data[i], heap.data[i/2] = heap.data[i/2], heap.data[i]
		i = i / 2
	}
}

//delete node in heap

func main() {
	h := New(30)
	for i := 0; i < 20; i++ {
		value := rand.Intn(20)
		fmt.Printf("%d ", value)
		h.Insert(value)
	}
	fmt.Println()
	for i := 1; i < 21; i++ {
		fmt.Printf("%d ", h.data[i])
	}
}

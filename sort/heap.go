package main

import (
	"fmt"
	"math/rand"
)

/*堆是一种特殊的树，它满足需要满足两个条件：
（1）堆是一种完全二叉树，也就是除了最后一层，其他层的节点个数都是满的，最后一个节点都靠左排列。
（2）堆中每一个节点的值都必须大于等于（或小于等于）其左右子节点的值。
对于每个节点的值都大于等于子树中每个节点值的堆，我们叫作“大顶堆”。对于每个节点的值都小于等于子树中每个节点值的堆，我们叫作“小顶堆”。

用数组来存储完全二叉树是非常节省内存空间的，因为我们不需要存储左右子节点的指针，单纯通过数组的下标，就可以找到一个节点的子节点和父节点。
数组中下标为 i 的节点的左子节点，就是下标为 i∗2 的节点，右子节点就是下标为 i∗2+1的节点，父节点就是下标为i/2的节点
*/

//https://blog.csdn.net/qq_55624813/article/details/121316293

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
	//从下往上堆化
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

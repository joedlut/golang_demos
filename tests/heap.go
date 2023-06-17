package main

import (
	"fmt"
	"math/rand"
)

func swap1(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func insert(arr []int, count int, value int) {
	fmt.Printf("count=%d\n", count)
	arr[count] = value
	n := count
	for n/2 >= 1 && arr[n/2] < arr[n] {
		swap1(arr, n/2, n)
		n /= 2
	}
}

func removeMax(arr []int, count int) {

}

func BuildHeap1(arr []int, capacity int) []int {
	heap := make([]int, capacity)
	//至少两个元素  arr[0] 0  arr[1] 第一个元素
	if len(arr) == 1 {
		return heap
	}
	// 需要有一个值表示现在堆中含有多少个元素
	for i := 1; i < len(arr); i++ {
		insert(heap, i, arr[i])
		fmt.Printf("heap -----%v\n", heap)
	}
	return heap
}

func heapAdjust(arr []int, n int, i int) {
	for {
		maxPos := i
		if n >= 2*i && arr[i] < arr[2*i] {
			maxPos = 2 * i
		}
		if n >= 2*i+1 && arr[maxPos] < arr[2*i+1] {
			maxPos = 2*i + 1
		}
		if maxPos == i {
			break
		}
		//不要忘了交换节点的值，大顶堆
		swap1(arr, maxPos, i)
		i = maxPos
	}
}

// n 堆元素的个数
func heapSort(arr []int, n int) {
	//建立初始堆的方式变了
	heap := BuildHeap1(arr, n+1)
	//fmt.Println(arr)
	k := n
	for k >= 1 {
		// 第一个节点跟 最后一个节点交换位置
		swap1(heap, 1, k)
		k--
		//从上往下 调整
		heapAdjust(heap, k, 1)
	}
	//输出数组的排序方式
	fmt.Println(heap)
}

func main() {
	arr := make([]int, 1)
	for i := 1; i <= 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	//heap := BuildHeap1(arr, len(arr))
	//fmt.Println(heap)
	heapSort(arr, len(arr))
}

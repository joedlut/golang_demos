package main

import (
	"fmt"
	"math/rand"
)

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func HeapOnce(arr []int, n int, i int) {
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
		swap(arr, maxPos, i)
		i = maxPos
	}
}

func BuildHeap(arr []int, n int) {
	for k := n / 2; k >= 1; k-- {
		HeapOnce(arr, n, k)
	}
}

func HeapSort(arr []int, n int) {
	BuildHeap(arr, n)
	//fmt.Println(arr)
	for k := n; k >= 1; {
		swap(arr, k, 1)
		k--
		HeapOnce(arr, k, 1)
	}
}

func main() {
	arr := make([]int, 1)
	for i := 1; i <= 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	HeapSort(arr, 10)
	fmt.Println(arr)
}

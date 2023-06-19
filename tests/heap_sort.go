package main

import (
	"fmt"
	"math/rand"
)

//交换两个数组结点

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//调堆， 从上往下调整，n 数组的个数，i 从哪个非叶子结点调整，从上往下调整

func HeapOnce(arr []int, n int, i int) {
	for {
		maxPos := i
		if n >= 2*i && arr[i] < arr[2*i] {
			maxPos = 2 * i
		}
		//注意，是和arr[maxPos]比较
		if n >= 2*i+1 && arr[maxPos] < arr[2*i+1] {
			maxPos = 2*i + 1
		}
		//退出循环的条件
		if maxPos == i {
			break
		}

		//不要忘了交换节点的值，大顶堆
		swap(arr, maxPos, i)
		i = maxPos
	}
}

//从第一个非叶子结点 n / 2 调整

func BuildHeap(arr []int, n int) {
	for k := n / 2; k >= 1; k-- {
		HeapOnce(arr, n, k)
	}
}

func HeapSort(arr []int, n int) {
	//初始建堆

	BuildHeap(arr, n)
	//fmt.Println(arr)
	for k := n; k >= 1; {
		//第一个结点跟最后一个结点交换
		swap(arr, k, 1)
		k--
		//从堆顶往下调

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

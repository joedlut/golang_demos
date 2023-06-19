package main

import "fmt"

//调堆

func HeapOnce(arr []int, n int, i int) {
	for {
		maxPos := i
		fmt.Println(i, maxPos)
		if 2*i <= n && arr[2*i] >= arr[maxPos] {
			maxPos = 2 * i
		}
		if 2*i+1 <= n && arr[2*i+1] >= arr[maxPos] {
			maxPos = 2*i + 1
		}
		if maxPos == i {
			break
		}
		arr[maxPos], arr[i] = arr[i], arr[maxPos]
		i = maxPos
	}
	fmt.Println("break")
}

//建堆

func BuildHeap(arr []int, n int) {
	for i := n / 2; i >= 1; i-- {
		HeapOnce(arr, n, i)
	}
}

//堆排序，先建好堆，然后每次最后一个元素跟堆顶交换，排除最后一个元素，从堆顶往下调整

func HeapSort(arr []int, n int) {
	BuildHeap(arr, n)
	for k := n; k >= 1; {
		arr[1], arr[k] = arr[k], arr[1]
		k--
		//交换后 去掉最后一个元素， 从堆顶往下调
		HeapOnce(arr, k, 1)
	}
}

func main() {
	arr := []int{0, 3, 2, 1, 12, 45, 101, 1000, 999}
	HeapSort(arr, 8)
	fmt.Println(arr)
}

package main

import "fmt"

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 建堆
// 从后往前处理数组，并且每个数据都是  从上往下  堆化。也就是从二叉树中第一个非叶子节点开始开始堆化，因为叶子节点往下堆化只能自己跟自己比较
func buildHeap(arr []int, n int) {
	for i := n / 2; i >= 1; i-- {
		heapOnce(arr, n, i)
	}
}

func heapOnce(arr []int, n, i int) {
	for {
		// maxPos = 1  第一次循环
		// maxPos = 3 第二次循环
		maxPos := i
		//有左子节点
		if n >= 2*i && arr[i] <= arr[2*i] {
			maxPos = 2 * i
		}
		//有右子节点
		////if n >= 2*i+1 && arr[i] <= arr[2*i+1] {
		if n >= 2*i+1 && arr[maxPos] <= arr[2*i+1] {
			maxPos = 2*i + 1
		}
		//第二次循环 3=3 退出循环
		if maxPos == i {
			break
		}
		swap(arr, i, maxPos)
		// i = 3  第一次循环
		i = maxPos
	}
}

func heapSort(arr []int, n int) {
	//建堆
	buildHeap(arr, n)
	fmt.Println(arr)
	k := n
	for k >= 1 {
		swap(arr, 1, k)
		k--
		//从上往下 调整
		heapOnce(arr, k, 1)
	}
}

func main() {
	arr := []int{0, 20, 2, 45, 6, 88}
	heapSort(arr, 5)
	fmt.Println(arr)
}

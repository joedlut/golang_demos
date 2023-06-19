package main

import "fmt"

//快速排序

func quickSort(arr []int, left int, right int) {
	temp := arr[left]
	p := left
	i, j := left, right
	for i < j {
		//这里两个与条件的顺序很重要，搞反了数组就会溢出
		for j >= p && temp < arr[j] {
			j--
		}
		if j >= p {
			arr[p] = arr[j]
			p = j
		}
		for i <= p && temp > arr[i] {
			i++
		}
		if i <= p {
			arr[p] = arr[i]
			p = i
		}
	}
	arr[p] = temp
	if p-left > 1 {
		quickSort(arr, left, p-1)
	}
	if right-p > 1 {
		quickSort(arr, p+1, right)
	}
}

func main() {
	arr := []int{3, 4, 5, 6, 7, 12, 345, -1, 11, 32, 9}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

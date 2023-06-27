package main

import "fmt"

/*
我们将输入数组data分解成4部分, 一个flag, 一个乱序数组A, 两个空的容器B1, B2. 这里需要脑补.
flag = 6 >A中的元素4, 因此要扔到容器B1中.
同理, flag = 6 >A中的元素2, 因此要扔到容器B1中.
flag = 6 <A中的元素8, 因此要扔到容器B2中. 这里有个小玄机, 将A的首尾元素交换, 同时更新B2, 就可以了.
与2同理.
A中元素已经是空的了, partition结束. 递归调用B1, B2.
*/

func qsort(arr []int, start, end int) {
	//只有0个元素 或者 1个元素  不用排序
	// 0个元素  0 0
	// 1个元素  0 1
	if end-start <= 1 {
		return
	}
	b1Start, b1End := start, start
	b2Start, b2End := end, end
	aStart, aEnd := start+1, end
	for aEnd-aStart > 0 {
		if arr[aStart-1] > arr[aStart] {
			arr[aStart-1], arr[aStart] = arr[aStart], arr[aStart-1]
			b1End++
			aStart++
		} else {
			//为什么是  aEnd - 1
			arr[aStart], arr[aEnd-1] = arr[aEnd-1], arr[aStart]
			b2Start--
			aEnd--
		}
	}
	qsort(arr, b1Start, b1End)
	qsort(arr, b2Start, b2End)
}

func main() {
	arr := []int{9, 80, 7, 6, 2, -1}
	qsort(arr, 0, 6)
	fmt.Println(arr)
}

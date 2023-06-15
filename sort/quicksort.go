package main

import "fmt"

/*
一趟快速排序的算法是： [1]
1）设置两个变量i、j，排序开始的时候：i=0，j=N-1； [1]
2）以第一个数组元素作为关键数据，赋值给key，即key=A[0]； [1]
3）从j开始向前搜索，即由后开始向前搜索(j--)，找到第一个小于key的值A[j]，将A[j]和A[i]的值交换
4）从i开始向后搜索，即由前开始向后搜索(i++)，找到第一个大于key的A[i]，将A[i]和A[j]的值交换；
5）重复第3、4步，直到i==j； (3,4步中，没找到符合条件的值，即3中A[j]不小于key,4中A[i]不大于key的时候改变j、i的值，
使得j=j-1，i=i+1，直至找到为止。找到符合条件的值，进行交换的时候i， j指针位置不变。另外，i==j这一过程一定正好是i+或j-完成的时候，此时令循环结束）。
*/

func quickSort(values []int, left int, right int) {
	// 以第一个数组元素作为关键数据，赋值给key，即key=A[0]  [1]
	temp := values[left]
	p := left
	i, j := left, right

	// {8, 2, 30, 4, 51, 23, 45, 78, 10, 9, 1, 12}
	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			// ????
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	//确定分割点 p 的位置
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func main() {
	arr := []int{1, 2, 30, 4, 51, 23, 45, 78, 10, 9, 23, 12}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

package main

import (
	"fmt"
	"math/rand"
)

//可以不用数组， 直接概率计算
/*func getNext(m map[string]int) string {
	sum := 0
	for _, weight := range m {
		sum += weight
	}
	arr := make([]string, 0)
	for key, value := range m {
		for i := 0; i < value; i++ {
			arr = append(arr, key)
		}
	}
	index := rand.Intn(sum)
	return arr[index]
}*/

/*
	func getNext(m map[string]int) string {
		sum := 0
		weightList := make([]int, 0)
		keyList := make([]string, 0)
		for key, value := range m {
			sum += value
			weightList = append(weightList, value)
			keyList = append(keyList, key)
		}
		index := rand.Intn(sum)
		l := len(weightList)
		var curWeightSum int
		var lastWeightSum int
		for i := 0; i < l; i++ {
			if i == 0 {
				lastWeightSum = 0
			}
			curWeightSum += weightList[i]
			if index >= lastWeightSum && index < curWeightSum {
				return keyList[i]
			}
			lastWeightSum = curWeightSum
		}
		return ""
	}
*/
func getNext(m map[string]int) string {
	sum := 0
	weightList := make([]int, 0)
	keyList := make([]string, 0)
	for key, value := range m {
		sum += value
		weightList = append(weightList, sum)
		keyList = append(keyList, key)
	}
	index := rand.Intn(sum)
	l := len(weightList)
	small := 0
	for i := 0; i < l; i++ {
		if i == 0 {
			small = 0
		} else {
			small = weightList[i-1]
		}

		if index >= small && index < weightList[i] {
			return keyList[i]
		}
	}
	return ""
}

func main() {
	m := make(map[string]int)
	m["a"] = 2
	m["b"] = 6
	m["c"] = 10
	count := map[string]int{}
	for i := 0; i < 9000; i++ {
		count[getNext(m)]++
	}
	fmt.Println(count)
}

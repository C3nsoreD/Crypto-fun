package main

import "fmt"

// Given an array of positive numbers and a positive # k fund the max sum of any contigouds subarry

var testArr = []int{2, 1, 5, 1, 3, 2}

func main() {
	fmt.Println(kthSum(testArr, 3))
}

func kthSum(arr []int, k int) int {
	sum := 0
	result := 0
	start := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i >= k-1 {
			result = max(result, sum)
			sum -= arr[start]
			start++
		}

	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

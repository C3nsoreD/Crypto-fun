// length of smallest subarry whose sum is less than S
package main

import (
	"fmt"
	"math/bits"
)

var testArry = []int{2, 1, 5, 2, 3, 2}
var testArry2 = []int{2, 1, 5, 2, 8}
var testArry3 = []int{3, 4, 1, 1, 6}

var MaxInt = (1<<bits.UintSize)/2 - 1

func main() {
	fmt.Println(smallestSubArray(testArry, 7))
	// fmt.Println(smallestSubArray(testArry2, 7))
	// fmt.Println(smallestSubArray(testArry3, 8))
}

func smallestSubArray(arr []int, s int) int {
	length := MaxInt
	sum := 0
	start := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		fmt.Println("outer")
		for sum >= s { // here's the trick
			fmt.Println("innder")
			length = min(length, i-start+1)
			sum -= arr[start]
			start++
		}
	}
	return length
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

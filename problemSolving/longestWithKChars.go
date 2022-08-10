// longest substring with maximum K distinct chars
package main

import (
	"fmt"
)

var testStr = "araaci"

func main() {
	// fmt.Println(longestKthDistinct(testStr, 1))
	fmt.Println(solution(testStr, 2))
	// fmt.Println(testStr[1:])
}

func longestKthDistinct(str string, k int) int {
	popCount := make(map[string]int)
	set := make(map[string]bool)
	length := 0
	distinct := 0
	strContainer := []string{}

	for _, v := range str {
		popCount[string(v)]++

	}

	fmt.Println(popCount)
	for _, v := range str {
		if _, ok := set[string(v)]; !ok {
			set[string(v)] = true
		}
		if distinct < k || popCount[string(v)] > 1 {
			strContainer = append(strContainer, string(v))
			length = max(length, len(strContainer))

			if set[string(v)] {
				distinct++
			}
			fmt.Println("distinct: ", distinct)
			fmt.Println("length: ", length)
			fmt.Println("strContainer: ", strContainer)
		}
	}
	return length
}

func solution(str string, k int) int {
	start := 0
	maxLength := 0
	popCount := make(map[string]int)
	var rC, lC string

	for i := 0; i < len(str); i++ {
		rC = string(str[i])
		if _, ok := popCount[string(rC)]; !ok {
			popCount[string(rC)] = 0
		}
		popCount[string(rC)]++

		for len(popCount) > k {
			lC = string(str[start])
			popCount[lC]--
			if popCount[lC] == 0 {
				delete(popCount, lC)
			}
			start++
		}
		maxLength = max(maxLength, i-start+1)
	}
	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*

Input: String="araaci", K=2
Output: 4
Explanation: The longest substring with no more than '2' distinct characters is "araa".

Input: String="araaci", K=1
Output: 2
Explanation: The longest substring with no more than '1' distinct characters is "aa".

*/

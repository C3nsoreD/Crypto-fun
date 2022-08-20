package main

var testArray = [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}

var testArrayTwo = [][]int{{5}}

// func main() {
// 	fmt.Println(kthSmallest(testArrayTwo, 1))
// }

func kthSmallest(matrix [][]int, k int) int {
	c := 1
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix[1]); j++ {
			if c == k {
				return matrix[i][j]
			}
			c++
		}
	}
	return -1
}

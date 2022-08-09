package main

import "fmt"

// Given an arr, find the average of all subarrays of K's contigous elements in it.

func main() {
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5
	averages := []float32{}
	for i := 0; i < len(arr); i++ {
		border := k + i - 1
		if border > len(arr)-1 {
			break
		}
		avg := average(arr[i:border], k)
		averages = append(averages, avg)
	}
	fmt.Println(averages)
}

func average(subArray []int, k int) float32 {
	sum := 0
	for _, v := range subArray {
		sum += v
	}
	// fmt.Println(sum / k)
	return float32(sum) / float32(k)
}

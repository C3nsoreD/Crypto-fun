package main

import "fmt"

// Given an arr, find the average of all subarrays of K's contigous elements in it.

func main() {
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 2
	// v1 := kthContigous(arr, k)

	v2 := slidingAvg(arr, k)
	// fmt.Println(v1)
	fmt.Println(v2)
}

func average(subArray []int, k int) float32 {
	sum := 0
	for _, v := range subArray {
		sum += v
	}
	return float32(sum) / float32(k)
}

func kthContigous(arr []int, k int) []float32 {
	averages := []float32{}
	for i := 0; i < len(arr); i++ {
		border := k + i
		if border > len(arr) {
			break
		}
		// fmt.Println(i, border, arr[i:border])
		avg := average(arr[i:border], k)
		averages = append(averages, avg)
	}
	return averages
}

func slidingAvg(arr []int, k int) []float32 {
	sum := 0
	start := 0
	v2Result := []float32{}
	// using sliding window
	for boarder := 0; boarder < len(arr); boarder++ {
		sum += arr[boarder]
		if boarder >= k-1 {
			v2Result = append(v2Result, float32(sum)/float32(k)) // store result
			sum -= arr[start]                                    // reduce current sum, worm like movement skills!!
			start++
		}
	}
	return v2Result
}

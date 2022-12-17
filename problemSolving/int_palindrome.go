// Given an integer x, return true if x is palindrome integer.
package main

// first solution
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    arr := numToArr(x)
    for i := range arr {
        if arr[i] != arr[len(arr)-1-i] {
            return false
        }
    }
    return true
}

func numToArr(x int) (arr []int) {
    for x > 0 {
        arr = append(arr, x%10)
        x = x / 10
    }
    return arr
}

func isPalindrome_two(x int) bool {
    var reverseNum int
    tmp := x
    for tmp > 0 {
        reverseNum = reverseNum*10 + tmp%10 // reconstruct
        tmp = tmp / 10
    }
    return x == reverseNum
}

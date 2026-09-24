package main

import "fmt"

// RunningSum returns a slice where element i is the sum of nums[0] through nums[i].
func RunningSum(nums []int) []int {
	result := make([]int, len(nums))
	sum := 0
	for i, n := range nums {
		sum += n
		result[i] = sum
	}
	return result
}

func main() {
	fmt.Println(RunningSum([]int{1, 2, 3, 4}))
}

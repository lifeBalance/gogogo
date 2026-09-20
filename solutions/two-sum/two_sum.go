package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	// Maps number value to its index in the slice
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil
}

func main() {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
	fmt.Println(TwoSum([]int{3, 2, 4}, 6))       // Output: [1, 2]
	fmt.Println(TwoSum([]int{3, 3}, 6))          // Output: [0, 1]
}
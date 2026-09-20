# Running Sum of 1d Array: editorial

## Approach

This is the simplest prefix sum. The running sum at index `i` is the running sum at index `i - 1` plus `nums[i]`. Keep a single accumulator and append it after each addition.

## Complexity

- Time: O(n).
- Space: O(1) beyond the output slice.

## Solution

```go
func RunningSum(nums []int) []int {
	out := make([]int, 0, len(nums))
	total := 0
	for _, x := range nums {
		total += x
		out = append(out, total)
	}
	return out
}
```

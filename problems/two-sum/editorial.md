# Two Sum: editorial

## Approach

Walk the slice once and keep a map from each value seen so far to its index. For the current number `x`, the partner we need is `target - x`. If it is already in the map, we are done. Otherwise, record `x` and move on.

Checking the map before inserting the current number is what prevents using the same element twice.

## Complexity

- Time: O(n). One pass, constant-time map operations.
- Space: O(n) for the map in the worst case.

## Solution

```go
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, x := range nums {
		if j, ok := seen[target-x]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return nil
}
```

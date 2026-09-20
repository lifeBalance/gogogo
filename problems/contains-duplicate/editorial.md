# Contains Duplicate: editorial

## Approach

A set holds each value once. Go has no set type, so use a `map[int]struct{}`. Walk the slice and add each element. If an element is already present, there is a duplicate and we can stop early.

## Complexity

- Time: O(n).
- Space: O(n) for the map.

## Solution

```go
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, x := range nums {
		if _, ok := seen[x]; ok {
			return true
		}
		seen[x] = struct{}{}
	}
	return false
}
```

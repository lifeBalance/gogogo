---
title: Running Sum of 1d Array
# 1 to 5 bananas. 1-2 easy, 3 medium, 4-5 hard. Leave it out for no rating.
difficulty: 2
tags: [array, prefix-sum]
starter: running_sum.go
---

Given a list of integers `nums`, return a list where the element at index `i` is the sum of `nums[0]` through `nums[i]`.

## Examples

```text
nums = [1, 2, 3, 4]      ->  [1, 3, 6, 10]
nums = [1, 1, 1, 1, 1]   ->  [1, 2, 3, 4, 5]
nums = [3, 1, 2, 10, 1]  ->  [3, 4, 6, 16, 17]
```

## Constraints

- `1 <= len(nums) <= 1000`
- `-10^6 <= nums[i] <= 10^6`

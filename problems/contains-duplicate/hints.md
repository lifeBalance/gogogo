1. Comparing every pair is quadratic. Is there a structure that tells you in constant time whether you have seen a value before?
2. A set stores values without duplicates. What happens to its size if you add every element of the list?
3. Either compare `len(set(nums))` with `len(nums)`, or walk the list adding to a set and stop the first time a value is already there.

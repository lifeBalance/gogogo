1. The brute force checks every pair. That works, but it is quadratic. What information would let you skip the inner loop?
2. For each number `x`, the number you are looking for is `target - x`. Can you check whether you have already seen it in constant time?
3. Walk the list once. Keep a dictionary from value to index. Before inserting `x`, look up `target - x`.

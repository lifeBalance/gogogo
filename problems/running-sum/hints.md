1. Recomputing the sum from the start for every index is quadratic. How does the answer at index `i` relate to the answer at index `i - 1`?
2. Each output element is the previous output element plus the current input element.
3. Keep one running total. For each `x`, add it to the total and append the total to the result.

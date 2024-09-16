# Problem Statement

We are given an unsorted array containing `n+1` numbers taken from the range `1` to `n`. The array has only one duplicate, but it can be repeated multiple times. **Find that duplicate number without using any extra space**. You are, however, allowed to modify the input array.

### Example 1:
```
Input: [1, 4, 4, 3, 2]
Output: 4
```

### Example 2:
```
Input: [2, 1, 3, 3, 5, 4]
Output: 3
```

### Example 3:
```
Input: [2, 4, 1, 4, 4]
Output: 4
```

## Constraints:

- `nums.length == n + 1`
- `1 <= n <= 10^5`
- `1 <= nums[i] <= n`
- All the integers in `nums` appear only once except for precisely one integer which appears two or more times.
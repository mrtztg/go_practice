# Intervals Intersection 

## Problem Statement

Given two lists of intervals, find the intersection of these two lists. Each list consists of disjoint intervals sorted on their start time.

### Example 1:

**Input:** `arr1=[[1, 3], [5, 6], [7, 9]], arr2=[[2, 3], [5, 7]]`

**Output:** `[2, 3], [5, 6], [7, 7]`

**Explanation:** The output list contains the common intervals between the two lists.

### Example 2:

**Input:** `arr1=[[1, 3], [5, 7], [9, 12]], arr2=[[5, 10]]`

**Output:** `[5, 7], [9, 10]`

**Explanation:** The output list contains the common intervals between the two lists.

## Constraints:

- `0 <= arr1.length, arr2.length <= 1000`
- `arr1.length + arr2.length >= 1`
- `0 <= starti < endi <= 10^9`
- `endi < starti + 1`
- `0 <= startj < endj <= 10^9`
- `endj < startj + 1`

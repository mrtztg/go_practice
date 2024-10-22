## Problem Statement

Given the head of a LinkedList and a number `k`, reverse every `k` sized sub-list starting from the head.

If, in the end, you are left with a sub-list with less than `k` elements, reverse it too.

### Example

Original List:

`1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> null`

k = 3

Reversed Sub-list:

`3 -> 2 -> 1 -> 6 -> 5 -> 4 -> 8 -> 7 -> null`

### Constraints:

- The number of nodes in the list is n.
- `1 <= k <= n <= 5000`
- `0 <= Node.val <= 1000`

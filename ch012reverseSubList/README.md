# Reverse a Sub-list (medium)

## Problem Statement

Given the head of a LinkedList and two positions `p` and `q`, reverse the LinkedList from position `p` to `q`.

### Example:

Original List:
```
1 -> 2 -> 3 -> 4 -> 5 -> null
```
p = 2, q = 4

Reversed Sub-list:
```
1 -> 4 -> 3 -> 2 -> 5 -> null
```

## Constraints:

- The number of nodes in the list is `n`.
- `1 <= n <= 500`
- `-500 <= Node.val <= 500`
- `1 <= p <= q <= n`
# Problem Statement

Given an array of intervals representing ‘N’ appointments, find out if a person can attend all the appointments.

## Example 1:

- **Appointments:** `[[1,4], [2,5], [7,9]]`
- **Output:** `false`
- **Explanation:** Since `[1,4]` and `[2,5]` overlap, a person cannot attend both of these appointments.

## Example 2:

- **Appointments:** `[[6,7], [2,4], [13, 14], [8,12], [45, 47]]`
- **Output:** `true`
- **Explanation:** None of the appointments overlap, therefore a person can attend all of them.

## Example 3:

- **Appointments:** `[[4,5], [2,3], [3,6]]`
- **Output:** `false`
- **Explanation:** Since `[4,5]` and `[3,6]` overlap, a person cannot attend both of these appointments.

## Constraints:

- `1 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= starti < endi <= 10^6`

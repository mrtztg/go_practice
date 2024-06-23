package ch001MergeIntervals

import (
	"cmp"
	"slices"
)

type Interval struct {
	Start, End int
}
type solution struct {
}

func (s *solution) merge(intervals []Interval) []Interval {
	slices.SortStableFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.Start, b.Start)
	})
	var merged []Interval
	for _, in := range intervals {
		if len(merged) == 0 {
			merged = append(merged, in)
			continue
		}
		if in.Start <= merged[len(merged)-1].End {
			merged[len(merged)-1].End = max(in.End, merged[len(merged)-1].End)
		} else {
			merged = append(merged, in)
		}
	}
	return merged
}

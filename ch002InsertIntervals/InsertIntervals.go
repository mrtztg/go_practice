package ch002InsertIntervals

type Interval struct {
	Start, End int
}

type solution struct {
}

func (s *solution) insert(intervals []Interval, newInterval Interval) []Interval {
	newIntervalAdded := false
	merged := make([]Interval, 0)
	a := func(interv Interval) {
		if len(merged) == 0 || merged[len(merged)-1].End < interv.Start {
			merged = append(merged, interv)
		} else {
			merged[len(merged)-1].End = max(merged[len(merged)-1].End, interv.End)
		}
	}
	for _, interval := range intervals {
		if !newIntervalAdded {
			if newInterval.Start < interval.Start {
				a(newInterval)
				newIntervalAdded = true
			}
		}
		a(interval)
	}
	if !newIntervalAdded {
		a(newInterval)
	}
	return merged
}

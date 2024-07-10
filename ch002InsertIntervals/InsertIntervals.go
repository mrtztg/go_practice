package ch002InsertIntervals

type Interval struct {
	Start, End int
}

type solution struct {
}

//func (s *solution) insert(intervals []Interval, newInterval Interval) []Interval {
//	newIntervalAdded := false
//	merged := make([]Interval, 0)
//	a := func(interv Interval) {
//		if len(merged) == 0 || merged[len(merged)-1].End < interv.Start {
//			merged = append(merged, interv)
//		} else {
//			merged[len(merged)-1].End = max(merged[len(merged)-1].End, interv.End)
//		}
//	}
//	for _, interval := range intervals {
//		if !newIntervalAdded {
//			if newInterval.Start < interval.Start {
//				a(newInterval)
//				newIntervalAdded = true
//			}
//		}
//		a(interval)
//	}
//	if !newIntervalAdded {
//		a(newInterval)
//	}
//	return merged
//}

func (s *solution) insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}
	merged := make([]Interval, 0)
	i := 0
	for i < len(intervals) && intervals[i].End < newInterval.Start {
		merged = append(merged, intervals[i])
		i++
	}
	for i < len(intervals) && newInterval.End >= intervals[i].Start {
		newInterval.Start = min(newInterval.Start, intervals[i].Start)
		newInterval.End = max(newInterval.End, intervals[i].End)
		i++
	}
	merged = append(merged, newInterval)
	for i < len(intervals) {
		merged = append(merged, intervals[i])
		i++
	}
	return merged
}

func min(val1, val2 int) int {
	if val1 > val2 {
		return val2
	}
	return val1
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

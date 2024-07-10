package ch003IntervalIntersection

type Interval struct {
	Start, End int
}

type solution struct {
}

func (s *solution) merge(arr1, arr2 []Interval) []Interval {
	i := 0
	j := 0
	merged := make([]Interval, 0)
	for i < len(arr1) && j < len(arr2) {
		intersectionInterval := Interval{
			Start: max(arr1[i].Start, arr2[j].Start),
			End:   min(arr1[i].End, arr2[j].End),
		}
		if intersectionInterval.End >= intersectionInterval.Start {
			merged = append(merged, intersectionInterval)
		}
		if intersectionInterval.End < intersectionInterval.Start && arr1[i].Start < arr2[j].Start || arr1[i].End < arr2[j].End {
			i++
		} else {
			j++
		}
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

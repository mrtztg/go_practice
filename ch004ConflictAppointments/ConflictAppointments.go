package ch004ConflictAppointments

import "sort"

type Interval struct {
	Start, End int
}

type Solution struct {
}

func (s *Solution) canAttendAllAppointments(intervals []Interval) bool {
	if len(intervals) <= 1 {
		return true
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			return false
		}
	}
	return true
}

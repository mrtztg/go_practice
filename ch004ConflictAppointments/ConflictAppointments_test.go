package ch004ConflictAppointments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanAttendAllAppointments(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		appointments []Interval
		output       bool
	}{
		{
			appointments: []Interval{{1, 4}, {2, 5}, {7, 9}},
			output:       false,
		}, {
			appointments: []Interval{{6, 7}, {2, 4}, {13, 14}, {8, 12}, {45, 47}},
			output:       true,
		}, {
			appointments: []Interval{{4, 5}, {2, 3}, {3, 6}},
			output:       false,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.output, s.canAttendAllAppointments(tc.appointments), "appointment: %+v", tc.appointments)
	}
}

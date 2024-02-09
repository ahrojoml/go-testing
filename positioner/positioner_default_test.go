package positioner_test

import (
	"fmt"
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLinearDistance(t *testing.T) {
	testCases := []struct {
		name     string
		from     *positioner.Position
		to       *positioner.Position
		expected float64
	}{
		{
			name: "all negative coordinates",
			from: &positioner.Position{
				X: -1,
				Y: -5,
				Z: -1,
			},
			to: &positioner.Position{
				X: -1,
				Y: -1,
				Z: -1,
			},
			expected: 4,
		}, {
			name: "all negative coordinates",
			from: &positioner.Position{
				X: 1,
				Y: 5,
				Z: 1,
			},
			to: &positioner.Position{
				X: 1,
				Y: 1,
				Z: 1,
			},
			expected: 4,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d, %s", idx, tc.name), func(t *testing.T) {
			positioner := positioner.NewPositionerDefault()

			linearDistance := positioner.GetLinearDistance(tc.from, tc.to)

			require.Equal(t, tc.expected, linearDistance)
		})
	}
}

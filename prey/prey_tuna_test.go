package prey_test

import (
	"fmt"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTunaPreyGetPosition(t *testing.T) {
	testCases := []struct {
		name     string
		position *positioner.Position
		expect   *positioner.Position
	}{
		{
			name:     "default position",
			position: &positioner.Position{},
			expect:   &positioner.Position{},
		}, {
			name: "non empty position",
			position: &positioner.Position{
				X: 1,
				Y: 1,
				Z: 1,
			},
			expect: &positioner.Position{
				X: 1,
				Y: 1,
				Z: 1,
			},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d, %s", idx, tc.name), func(t *testing.T) {
			tuna := prey.NewTuna(0, tc.position)

			result := tuna.GetPosition()

			require.Equal(t, tc.expect, result)
		})
	}
}

func TestTunaPreyGetSpeed(t *testing.T) {
	testCases := []struct {
		name   string
		speed  float64
		expect float64
	}{
		{
			name: "default speed",
		}, {
			name:   "non empty position",
			speed:  1,
			expect: 1,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d, %s", idx, tc.name), func(t *testing.T) {
			tuna := prey.NewTuna(tc.speed, &positioner.Position{})

			result := tuna.GetSpeed()

			require.Equal(t, tc.expect, result)
		})
	}
}

package simulator_test

import (
	"fmt"
	"testdoubles/positioner"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCatch(t *testing.T) {
	testCases := []struct {
		name     string
		hunter   *simulator.Subject
		prey     *simulator.Subject
		distance float64
		expect   bool
	}{
		{
			name: "all negative coordinates, hunter faster",
			hunter: &simulator.Subject{
				Position: &positioner.Position{
					X: -10,
					Y: -10,
					Z: -10,
				},
				Speed: 10,
			},
			prey: &simulator.Subject{
				Position: &positioner.Position{
					X: -5,
					Y: -5,
					Z: -5,
				},
				Speed: 5,
			},
			distance: 10,
			expect:   true,
		}, {
			name: "all negative coordinates, prey faster",
			hunter: &simulator.Subject{
				Position: &positioner.Position{
					X: -10,
					Y: -10,
					Z: -10,
				},
				Speed: 10,
			},
			prey: &simulator.Subject{
				Position: &positioner.Position{
					X: -5,
					Y: -5,
					Z: -5,
				},
				Speed: 15,
			},
			distance: 10,
			expect:   false,
		}, {
			name: "all positive coordinates, hunter faster",
			hunter: &simulator.Subject{
				Position: &positioner.Position{
					X: 10,
					Y: 10,
					Z: 10,
				},
				Speed: 10,
			},
			prey: &simulator.Subject{
				Position: &positioner.Position{
					X: 5,
					Y: 5,
					Z: 5,
				},
				Speed: 5,
			},
			distance: 10,
			expect:   true,
		}, {
			name: "all positive coordinates, prey faster",
			hunter: &simulator.Subject{
				Position: &positioner.Position{
					X: 10,
					Y: 10,
					Z: 10,
				},
				Speed: 10,
			},
			prey: &simulator.Subject{
				Position: &positioner.Position{
					X: -5,
					Y: -5,
					Z: -5,
				},
				Speed: 15,
			},
			distance: 10,
			expect:   false,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d, %s", idx, tc.name), func(t *testing.T) {
			ps := positioner.NewPositionerStub()
			ps.GetLinearDistanceFunc = func(hunter, prey *positioner.Position) float64 {
				return tc.distance
			}

			cs := simulator.NewCatchSimulatorDefault(100, ps)
			result := cs.CanCatch(tc.hunter, tc.prey)

			require.Equal(t, tc.expect, result)
		})
	}
}

func TestCanCatchIntegration(t *testing.T) {
	testCases := []struct {
		name   string
		hunter *simulator.Subject
		prey   *simulator.Subject
		expect bool
	}{
		{
			name: "hunter is faster",
			hunter: &simulator.Subject{
				Position: &positioner.Position{
					X: -10,
					Y: -10,
					Z: -10,
				},
				Speed: 10,
			},
			prey: &simulator.Subject{
				Position: &positioner.Position{
					X: -5,
					Y: -5,
					Z: -5,
				},
				Speed: 5,
			},
			expect: true,
		}, {
			name: "hunter is slower",
			hunter: &simulator.Subject{
				Position: &positioner.Position{
					X: -10,
					Y: -10,
					Z: -10,
				},
				Speed: 10,
			},
			prey: &simulator.Subject{
				Position: &positioner.Position{
					X: -5,
					Y: -5,
					Z: -5,
				},
				Speed: 15,
			},
			expect: false,
		}, {
			name: "hunter faster, not in time",
			hunter: &simulator.Subject{
				Position: &positioner.Position{
					X: 100,
					Y: 100,
					Z: 100,
				},
				Speed: 6,
			},
			prey: &simulator.Subject{
				Position: &positioner.Position{
					X: 5,
					Y: 5,
					Z: 5,
				},
				Speed: 5,
			},
			expect: false,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d, %s", idx, tc.name), func(t *testing.T) {
			sim := simulator.NewCatchSimulatorDefault(100, positioner.NewPositionerDefault())

			result := sim.CanCatch(tc.hunter, tc.prey)

			require.Equal(t, tc.expect, result)
		})
	}
}

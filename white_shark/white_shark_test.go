package hunt_test

import (
	"fmt"
	hunt "testdoubles/white_shark"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	testCases := []struct {
		name      string
		hungry    bool
		tired     bool
		speed     float64
		tunaSpeed float64
		expect    error
	}{
		{
			name:   "shark is tired",
			hungry: true,
			tired:  true,
			expect: hunt.ErrSharkIsTired,
		}, {
			name:   "shark ate and full",
			hungry: false,
			tired:  false,
			expect: hunt.ErrSharkIsNotHungry,
		}, {
			name:      "tuna is faster",
			hungry:    true,
			tired:     false,
			speed:     10,
			tunaSpeed: 20,
			expect:    hunt.ErrSharkIsSlower,
		}, {
			name:      "success",
			hungry:    true,
			tired:     false,
			speed:     20,
			tunaSpeed: 10,
			expect:    nil,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d: %s", idx, tc.name), func(t *testing.T) {
			shark := hunt.NewWhiteShark(tc.hungry, tc.tired, tc.speed)
			tuna := hunt.NewTuna("test", tc.tunaSpeed)

			err := shark.Hunt(tuna)

			if tc.expect != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expect)
			} else {
				require.NoError(t, err)
				require.False(t, shark.Hungry)
				require.True(t, shark.Tired)
			}
		})
	}
}

func TestWhiteSharkHunt_NilTuna(t *testing.T) {
	t.Run("tuna is nil", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 10)

		err := shark.Hunt(nil)

		require.Error(t, err)
		require.ErrorIs(t, err, hunt.ErrTunaIsNil)
	})

}

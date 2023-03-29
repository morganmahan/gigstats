package stats

import (
	"testing"
)

func TestCountUniqueElements(t *testing.T) {
	t.Run("Count unique elements", func(t *testing.T) {
		result := CountUniqueElements([]string{
			"ABC",
			"ABC",
			"DEF",
			"DEF",
			"def",
		})
		if result != 2 {
			t.Errorf("Unique elements count incorrect")
		}
	})
}

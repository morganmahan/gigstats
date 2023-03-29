package splitter

import (
	"reflect"
	"testing"
)

func TestSplitCommaSeparatedArrayValues(t *testing.T) {
	t.Run("Should split comma separated array values", func(t *testing.T) {
		expected := []string{
			"Value 1",
			"Value 2",
			"Value 3",
			"Value 4",
			"Value 5",
		}
		result := SplitCommaSeparatedArrayValues([]string{
			"Value 1, Value 2, Value 3",
			"Value 4, Value 5",
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Result does not match expected")
		}
	})
}

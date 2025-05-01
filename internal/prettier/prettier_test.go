package prettier

import "testing"

func TestStandardiseCommaSeparated(t *testing.T) {
	t.Run("", func(t *testing.T) {
		result := StandardiseCommaSeparated("Abc, abc, ABC")
		if result != "Abc, Abc, Abc" {
			t.Errorf("String not standardised when comma separated")
		}
	})
}

func TestMakeStringArrayCommaSeparatedString(t *testing.T) {
	t.Run("", func(t *testing.T) {
		result := MakeStringArrayCommaSeparatedString([]string{
			"1",
			"2",
			"3",
		})
		if result != "1, 2, 3" {
			t.Errorf("String array not correctly comma separated")
		}
	})
}

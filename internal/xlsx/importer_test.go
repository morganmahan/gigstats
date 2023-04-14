package xlsx

import (
	"reflect"
	"testing"
)

func TestGetColumns(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		expected := GigSheet{
			Bands: [][]string{
				{"Alter Bridge", "Black Stone Cherry"},
				{"Black Veil Brides"},
			},
			Venue: []string{
				"Wembley Arena",
				"Brixton Academy",
			},
			Date: []string{
				"29/11/2011",
				"30/03/2012",
			},
			Who: [][]string{
				{"Mum", "Alex"},
				{"Harry"},
			},
			Tour: []string{
				"Abiii",
				"Stwof",
			},
			Hotel: []string{
				"Novotel",
			},
		}
		result, err := GetColumns("../../fixtures/gigs.xlsx")
		if err != nil {
			t.Errorf("GetColumns Errored: %d", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GetColumns result does not match expected")
		}
	})
}

func TestGetSheetRows(t *testing.T) {
	t.Run("Return a 2D array of rows", func(t *testing.T) {
		expected := [][]string{
			{
				"Bands",
				"Venue",
				"Date",
				"Who",
				"Tour",
				"Hotel",
				"Significant Events",
			},
			{
				"Alter Bridge, Black Stone Cherry",
				"Wembley Arena",
				"29/11/2011",
				"Mum, Alex",
				"ABIII",
				"Novotel",
				"DVD Filmed",
			},
			{
				"Black Veil Brides",
				"Brixton Academy",
				"30/03/2012",
				"Harry",
				"STWOF",
				"",
				"Test",
			},
		}
		result, err := GetSheetRows("../../fixtures/gigs.xlsx")
		if err != nil {
			t.Errorf("getSheetRows errored")
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GetSheetRows returns incorrect response")
		}
	})
}

func TestGetColumnsByType(t *testing.T) {
	t.Run("Return a map of column names to an array of column cells", func(t *testing.T) {
		expected := GigSheet{
			Bands: [][]string{
				{"Test Band"},
			},
			Venue: []string{
				"Test Venue",
			},
			Date: []string{
				"29/11/2011",
			},
			Who: [][]string{
				{"Me"},
			},
			Tour: []string{
				"Test Tour",
			},
			Hotel: []string{
				"Test Hotel",
			},
		}
		result, err := getColumnsByType([][]string{
			{
				"Bands",
				"Venue",
				"Date",
				"Who",
				"Tour",
				"Hotel",
			},
			{
				"Test Band",
				"Test Venue",
				"29/11/2011",
				"Me",
				"Test Tour",
				"Test Hotel",
			},
		})
		if err != nil {
			t.Errorf("getColumnsByType errored")
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GetColumnsByType returns incorrect result")
		}
	})
}

func TestCheckSheetValidity(t *testing.T) {
	t.Run("Return true for a valid sheet", func(t *testing.T) {
		result := checkSheetValidity([][]string{
			{
				"Bands",
				"Venue",
				"Date",
				"Who",
				"Tour",
				"Hotel",
			},
		})
		if result != true {
			t.Errorf("Incorrectly marks a correctly formatted sheet as invalid")
		}
	})
	t.Run("Return false for an invalid sheet", func(t *testing.T) {
		result := checkSheetValidity([][]string{
			{
				"Hotel",
				"Venue",
				"Date",
				"Who",
				"Bands",
				"Tour",
			},
		})
		if result != false {
			t.Errorf("Incorrectly marks a badly formatted sheet as valid")
		}
	})
}

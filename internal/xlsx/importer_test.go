package xlsx

import (
	"reflect"
	"testing"
)

func TestGetColumns(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		expected := map[string][]string{
			"Bands": {
				"Alter Bridge, Black Stone Cherry",
				"Black Veil Brides",
			},
			"Venue": {
				"Wembley Arena",
				"Brixton Academy",
			},
			"Date": {
				"29/11/2011",
				"30/03/2012",
			},
			"Who": {
				"Mum, Alex",
				"Harry",
			},
			"Tour": {
				"ABIII",
				"STWOF",
			},
			"Hotel": {
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
		result, err := getSheetRows("../../fixtures/gigs.xlsx")
		if err != nil {
			t.Errorf("getSheetRows errored")
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Rows are incorrect")
		}
	})
}

func TestGetColumnsByType(t *testing.T) {
	t.Run("Return a map of column names to an array of column cells", func(t *testing.T) {
		expected := map[string][]string{
			"Bands": {
				"Test Band",
			},
			"Venue": {
				"Test Venue",
			},
			"Date": {
				"29/11/2011",
			},
			"Who": {
				"Me",
			},
			"Tour": {
				"Test Tour",
			},
			"Hotel": {
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
			t.Errorf("Columns are incorrect")
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

package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/morganmahan/gigstats/internal/xlsx"
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

func TestGetGigsForElement(t *testing.T) {
	t.Run("Return gig for an element occuring once, where its value is not in an array", func(t *testing.T) {
		result := getGigsForElement(xlsx.GigSheet{
			Bands: [][]string{
				{"Band1"},
				{"Band2"},
			},
			Venue: []string{
				"Venue1",
				"Venue2",
			},
			Date: []string{
				"Date1",
				"Date2",
			},
			Who: [][]string{
				{"Who1"},
				{"Who2"},
			},
		}, "Venue1", "Venue")
		if !reflect.DeepEqual(result, [][]string{
			{"Band1", "Venue1", "Date1", "Who1"},
		}) {
			fmt.Println(result)
			t.Errorf("Incorrect gigs returned for given element")
		}
	})
	t.Run("Return gig for an element occuring once, where its value is in an array", func(t *testing.T) {
		result := getGigsForElement(xlsx.GigSheet{
			Bands: [][]string{
				{"Band1", "Band3"},
				{"Band2"},
			},
			Venue: []string{
				"Venue1",
				"Venue2",
			},
			Date: []string{
				"Date1",
				"Date2",
			},
			Who: [][]string{
				{"Who1"},
				{"Who2"},
			},
		}, "Band1", "Bands")
		if !reflect.DeepEqual(result, [][]string{
			{"Band1, Band3", "Venue1", "Date1", "Who1"},
		}) {
			fmt.Println(result)
			t.Errorf("Incorrect gigs returned for given element")
		}
	})
	t.Run("Return gigs for an element occuring multiple times, where its value is not in an array", func(t *testing.T) {
		result := getGigsForElement(xlsx.GigSheet{
			Bands: [][]string{
				{"Band1"},
				{"Band2"},
				{"Band3"},
			},
			Venue: []string{
				"Venue1",
				"Venue2",
				"Venue2",
			},
			Date: []string{
				"Date1",
				"Date2",
				"Date3",
			},
			Who: [][]string{
				{"Who1"},
				{"Who2"},
				{"Who3"},
			},
		}, "Venue2", "Venue")
		if !reflect.DeepEqual(result, [][]string{
			{"Band2", "Venue2", "Date2", "Who2"},
			{"Band3", "Venue2", "Date3", "Who3"},
		}) {
			fmt.Println(result)
			t.Errorf("Incorrect gigs returned for given element")
		}
	})
	t.Run("Return gigs for an element occuring multiple times, where its value is in an array", func(t *testing.T) {
		result := getGigsForElement(xlsx.GigSheet{
			Bands: [][]string{
				{"Band1"},
				{"Band2", "Band4"},
				{"Band2"},
			},
			Venue: []string{
				"Venue1",
				"Venue2",
				"Venue3",
			},
			Date: []string{
				"Date1",
				"Date2",
				"Date3",
			},
			Who: [][]string{
				{"Who1"},
				{"Who2"},
				{"Who3"},
			},
		}, "Band2", "Bands")
		if !reflect.DeepEqual(result, [][]string{
			{"Band2, Band4", "Venue2", "Date2", "Who2"},
			{"Band2", "Venue3", "Date3", "Who3"},
		}) {
			fmt.Println(result)
			t.Errorf("Incorrect gigs returned for given element")
		}
	})
}

package main

import (
	"fmt"
	"os"

	"github.com/morganmahan/gigstats/internal/stats"
	"github.com/morganmahan/gigstats/internal/xlsx"
	"github.com/morganmahan/gigstats/pkg/prettier"
)

func main() {
	// get all column cells keyed by their column name
	columns, err := xlsx.GetColumns("gigs.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// call stats functions
	stat := os.Args[1]
	var argument string
	if len(os.Args) > 2 {
		argument = prettier.Standardise(os.Args[2])
	}
	switch stat {
	case "bandsseen":
		fmt.Println(stats.CountUniqueElements(prettier.FlattenArray(columns.Bands)))
	case "bandcounts":
		prettier.PrintKeyValueArray(stats.GetOccurences(prettier.FlattenArray(columns.Bands)))
	case "venuesattended":
		fmt.Println(stats.CountUniqueElements(columns.Venue))
	case "venuecounts":
		prettier.PrintKeyValueArray(stats.GetOccurences(columns.Venue))
	case "venuegigs":
		prettier.Print2DArray(stats.GetGigsForVenue(columns, argument))
	case "bandgigs":
		prettier.Print2DArray(stats.GetGigsForBand(columns, argument))
	case "persongigs":
		prettier.Print2DArray(stats.GetGigsForPerson(columns, argument))
	}
}

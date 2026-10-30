package main

import (
	"fmt"
	"os"

	"github.com/morganmahan/gigstats/internal/prettier"
	"github.com/morganmahan/gigstats/internal/stats"
	"github.com/morganmahan/gigstats/internal/xlsx"
)

func main() {
	rows, err := xlsx.GetRows("gigs.xlsx")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	if !xlsx.CheckSheetValidity(rows) {
		fmt.Println("invalid sheet provided")
		return
	}
	// get all column cells keyed by their column name
	columns, err := xlsx.GetColumns(rows)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	gigs, err := xlsx.GetGigs(rows)
	if err != nil {
		fmt.Printf("Error: %v", err)
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
	case "venuesattended":
		fmt.Println(stats.CountUniqueElements(columns.Venue))
	case "bandcounts":
		prettier.PrintKeyValueArray(stats.GetOccurences(prettier.FlattenArray(columns.Bands)))
	case "personcounts":
		prettier.PrintKeyValueArray(stats.GetOccurences(prettier.FlattenArray(columns.Who)))
	case "venuecounts":
		prettier.PrintKeyValueArray(stats.GetOccurences(columns.Venue))
	case "venuegigs":
		prettier.PrintGigsArray(stats.GetGigsForVenue(gigs, argument))
	case "bandgigs":
		prettier.PrintGigsArray(stats.GetGigsForBand(gigs, argument))
	case "persongigs":
		prettier.PrintGigsArray(stats.GetGigsForPerson(gigs, argument))
	}
}

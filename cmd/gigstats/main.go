package main

import (
	"fmt"
	"os"

	"github.com/morganmahan/gigstats/internal/stats"
	"github.com/morganmahan/gigstats/internal/xlsx"
	"github.com/morganmahan/gigstats/pkg/prettier"
	"github.com/morganmahan/gigstats/pkg/splitter"
)

func main() {
	// get all column cells keyed by their column name
	columns, err := xlsx.GetColumns("gigs.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// loop over columns, splitting strings by commas
	columns["Bands"] = splitter.SplitCommaSeparatedArrayValues(columns["Bands"])
	columns["Who"] = splitter.SplitCommaSeparatedArrayValues(columns["Who"])

	// call stats functions
	stat := os.Args[1]
	switch stat {
	case "bandsseen":
		fmt.Println(stats.CountUniqueElements(columns["Bands"]))
	case "bandcounts":
		prettier.PrintKeyValueArray(stats.GetOccurences(columns["Bands"]))
	}
}

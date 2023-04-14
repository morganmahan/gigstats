package stats

import (
	"reflect"
	"sort"

	"github.com/morganmahan/gigstats/internal/xlsx"
	"github.com/morganmahan/gigstats/pkg/prettier"
	"golang.org/x/exp/slices"
)

func CountUniqueElements(elements []string) int {
	occurences := getOccurencesAsMap(elements)
	return len(occurences)
}

func GetOccurences(elements []string) []prettier.KeyValue {
	occurences := []prettier.KeyValue{}
	for key, value := range getOccurencesAsMap(elements) {
		occurences = append(occurences, prettier.KeyValue{Key: key, Value: value})
	}
	sort.Slice(occurences, func(i, j int) bool {
		return occurences[i].Value > occurences[j].Value
	})
	return occurences
}

func GetGigsForBand(cols xlsx.GigSheet, band string) [][]string {
	return getGigsForElement(cols, band, "Bands")
}

func GetGigsForVenue(cols xlsx.GigSheet, venue string) [][]string {
	return getGigsForElement(cols, venue, "Venue")
}

func GetGigsForPerson(cols xlsx.GigSheet, person string) [][]string {
	return getGigsForElement(cols, person, "Who")
}

func getOccurencesAsMap(elements []string) map[string]int {
	seen := map[string]int{}
	for _, elem := range elements {
		seen[elem] += 1
	}
	return seen
}

func getGigsForElement(cols xlsx.GigSheet, element string, elementType string) [][]string {
	indexMap := map[string]int{
		"Bands": 0,
		"Venue": 1,
		"Date":  2,
		"Who":   3,
		"Tour":  4,
		"Hotel": 5,
	}
	gigs := [][]string{}
	columns := reflect.ValueOf(cols)
	elementTypeColumn := columns.Field(indexMap[elementType])

	// Loop over the values of the given elementType (e.g. over the Bands array)
	for i := 0; i < elementTypeColumn.Len(); i++ {
		currentElement := elementTypeColumn.Index(i).Interface()
		// The element is either a string or a string array
		currentElementType := reflect.TypeOf(currentElement).String()

		// If the current element matches the element we are searching for
		if currentElementType == "string" {
			if currentElement == element {
				gigs = append(gigs, getGigAtIndex(columns, i))
			}
		} else {
			stringArr := currentElement.([]string)
			if slices.Contains(stringArr, element) {
				gigs = append(gigs, getGigAtIndex(columns, i))
			}
		}
	}
	return gigs
}

func getGigAtIndex(columns reflect.Value, index int) []string {
	// Loop over each column, adding the field at the given index to the results array
	gig := []string{}
	for i := 0; i < columns.NumField() && i <= 3; i++ {
		column := columns.Field(i)
		if column.Len() > index {
			elementAtIndex := column.Index(index).Interface()
			elementAtIndexType := reflect.TypeOf(elementAtIndex).String()
			if elementAtIndexType == "string" {
				gig = append(gig, elementAtIndex.(string))
			} else {
				stringArr := elementAtIndex.([]string)
				gig = append(gig, prettier.MakeStringArrayCommaSeparatedString(stringArr))
			}
		}
	}
	return gig
}

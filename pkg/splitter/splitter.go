package splitter

import "strings"

func SplitCommaSeparatedArrayValues(valsToSplit []string) []string {
	valsSplitByComma := []string{}
	for _, value := range valsToSplit {
		valsSplitByComma = append(valsSplitByComma, strings.Split(value, ", ")...)
	}
	return valsSplitByComma
}

func FlattenArray(valsToSplit [][]string) []string {
	vals := []string{}
	for _, value := range valsToSplit {
		vals = append(vals, value...)
	}
	return vals
}

package splitter

import "strings"

func SplitCommaSeparatedArrayValues(valsToSplit []string) []string {
	valsSplitByComma := []string{}
	for _, value := range valsToSplit {
		valsSplitByComma = append(valsSplitByComma, strings.Split(value, ", ")...)
	}
	return valsSplitByComma
}

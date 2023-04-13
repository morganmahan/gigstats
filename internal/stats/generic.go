package stats

import (
	"sort"

	"github.com/morganmahan/gigstats/pkg/prettier"
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

func getOccurencesAsMap(elements []string) map[string]int {
	seen := map[string]int{}
	for _, elem := range elements {
		standardisedElem := prettier.Standardise(elem)
		seen[standardisedElem] += 1
	}
	return seen
}

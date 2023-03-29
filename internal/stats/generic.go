package stats

import "strings"

func CountUniqueElements(elements []string) int {
	seen := map[string]int{}
	for _, elem := range elements {
		standardisedElem := strings.ToLower(elem)
		if seen[standardisedElem] == 0 {
			seen[standardisedElem] = 1
		}
	}
	return len(seen)
}

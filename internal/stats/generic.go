package stats

import (
	"sort"

	"github.com/morganmahan/gigstats/internal/prettier"
	"github.com/morganmahan/gigstats/internal/types"
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

func GetGigsForBand(gigs []types.Gig, band string) []types.Gig {
	res := []types.Gig{}
	for _, gig := range gigs {
		if slices.Contains(gig.Bands, band) {
			res = append(res, gig)
		}
	}
	return res
}

func GetGigsForPerson(gigs []types.Gig, person string) []types.Gig {
	res := []types.Gig{}
	for _, gig := range gigs {
		if slices.Contains(gig.Who, person) {
			res = append(res, gig)
		}
	}
	return res
}

func GetGigsForVenue(gigs []types.Gig, venue string) []types.Gig {
	res := []types.Gig{}
	for _, gig := range gigs {
		if gig.Venue == venue {
			res = append(res, gig)
		}
	}
	return res
}

func getOccurencesAsMap(elements []string) map[string]int {
	seen := map[string]int{}
	for _, elem := range elements {
		seen[elem] += 1
	}
	return seen
}

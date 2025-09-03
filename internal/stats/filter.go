package stats

import (
	"slices"

	"github.com/morganmahan/gigstats/internal/types"
)

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

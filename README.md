# Gig Stats

A program to create stats from a XLSX spreadsheet I have kept over the years containing information about gigs/events that I have been to

## Running

To get the stat you wish, you must provide the [stat type](#stat-types) string for that stat when running the program:

```
go run cmd/gigstats/main.go {argument}
```

## Testing

```
go test ./...
```

## XLSX Format

A [template](./template.xlsx) for the input XLSX file for this program has been provided.
Examples of the content structure can be found in the [test fixture file](./fixtures/gigs.xlsx).

## Stat Types

Total number of different bands seen (each band counted only once)
- `bandsseen`

List all bands and number of times seen
- `bandcounts`

## Future Improvements / Stats

`stats.GetOccurences` should return a key of the band/venue/whatever element, and a value of an array of each gig that element was a part of.
From this, add stats for:
- A rundown of a specific bands gigs
- A rundown of a specific venues gigs
- A rundown of a specific persons gigs
- Bands seen at particular festivals

Other stats

- How many gigs seen with each person
- How many gigs at each venue
- How many gigs in each year
- How many gigs in each month irrespective of year
- How many times stayed at each hotel
- How many time been to each tour/festival

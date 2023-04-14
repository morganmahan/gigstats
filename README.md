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

Total number of different venues seen (each venue counted only once)
- `venuesattended`

List all venues and number of times attended
- `venuecounts`

List details of all gigs at a particular venue
- `venuegigs "venuename"`

List details of all gigs for a particular band
- `bandgigs "bandname"`

List details of all gigs seen with a particular person
- `persongigs "personname"`

## Future Improvements / Stats

`stats.getGigsForElement` should take in `[][]string` of rows, rather than working over the columns, as this would allow columns that aren't necessariy filled for each gig to be included in the results. Currently, as the missing values are removed in the columns arrays, the indexes of a particular gig with the index of, say, its corresponding tour name, will not line up.

Move commands from main.go to their own package

- Bands seen at particular festivals
- How many gigs seen with each person
- How many gigs at each venue
- How many gigs in each year
- How many gigs in each month irrespective of year
- How many times stayed at each hotel
- How many time been to each tour/festival

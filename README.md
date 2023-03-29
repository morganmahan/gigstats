# Gig Stats

A program to create stats from a XLSX spreadsheet I have kept over the years containing information about gigs/events that I have been to

## Running

```
go run cmd/gigstats/main.go
```

## Testing

```
go test ./...
```

## XLSX Format

A [template](./template.xlsx) for the input XLSX file for this program has been provided.
Examples of the content structure can be found in the [test fixture file](./fixtures/gigs.xlsx).

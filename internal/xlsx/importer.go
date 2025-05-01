package xlsx

import (
	"fmt"
	"strings"

	"github.com/morganmahan/gigstats/internal/prettier"
	"github.com/morganmahan/gigstats/internal/types"
	"github.com/xuri/excelize/v2"
)

type GigSheet struct {
	Bands [][]string
	Venue []string
	Date  []string
	Who   [][]string
	Tour  []string
	Hotel []string
}

func GetRows(path string) ([][]string, error) {
	sheet, err := excelize.OpenFile(path)
	if err != nil {
		return [][]string{}, err
	}
	defer func() {
		err := sheet.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := sheet.GetRows("Sheet1")
	return rows, err
}

func GetGigs(rows [][]string) ([]types.Gig, error) {
	gigs := make([]types.Gig, len(rows))
	for i, row := range rows {
		// first row is headers
		if i == 0 {
			continue
		}
		// TODO: make tour and hotel import into the correct columns if one is missing
		// tour := ""
		// if len(row) > 4 {
		// 	tour = row[4]
		// }
		// hotel := ""
		// if len(row) > 5 {
		// 	hotel = row[5]
		// }
		gig := types.Gig{
			Bands: getArrayFromCommaList(row[0]),
			Venue: row[1],
			Date:  row[2],
			Who:   getArrayFromCommaList(row[3]),
			// Tour:  tour,
			// Hotel: hotel,
		}
		gigs = append(gigs, gig)
	}
	return gigs, nil
}

func GetColumns(rows [][]string) (GigSheet, error) {
	cols, err := getColumnsByType(rows)
	if err != nil {
		return GigSheet{}, err
	}

	return cols, nil
}

func CheckSheetValidity(rows [][]string) bool {
	headingRow := rows[0]
	if headingRow[0] != "Bands" || headingRow[1] != "Venue" || headingRow[2] != "Date" ||
		headingRow[3] != "Who" || headingRow[4] != "Tour" || headingRow[5] != "Hotel" {
		return false
	}
	return true
}

func getArrayFromCommaList(s string) []string {
	return strings.Split(prettier.StandardiseCommaSeparated(s), ", ")
}

func getColumnsByType(rows [][]string) (GigSheet, error) {
	columnsByType := GigSheet{}
	for _, row := range rows[1:] {
		columnsByType = appendValuesFromRowToColumns(columnsByType, row)
	}
	return columnsByType, nil
}

func appendValuesFromRowToColumns(columnsByType GigSheet, row []string) GigSheet {
	if row[0] != "" {
		columnsByType.Bands = append(columnsByType.Bands, getArrayFromCommaList(row[0]))
	}

	if row[1] != "" {
		columnsByType.Venue = append(columnsByType.Venue, prettier.Standardise(row[1]))
	}

	if row[2] != "" {
		columnsByType.Date = append(columnsByType.Date, prettier.Standardise(row[2]))
	}

	if row[3] != "" {
		columnsByType.Who = append(columnsByType.Who, getArrayFromCommaList(row[3]))
	}

	if len(row) > 4 && row[4] != "" {
		columnsByType.Tour = append(columnsByType.Tour, prettier.Standardise(row[4]))
	}

	if len(row) > 5 && row[5] != "" {
		columnsByType.Hotel = append(columnsByType.Hotel, prettier.Standardise(row[5]))
	}

	return columnsByType
}

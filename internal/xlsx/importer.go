package xlsx

import (
	"fmt"
	"strings"

	"github.com/morganmahan/gigstats/pkg/prettier"
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

func GetColumns(path string) (GigSheet, error) {
	rows, err := GetSheetRows(path)
	if err != nil {
		return GigSheet{}, err
	}
	cols, err := getColumnsByType(rows)
	if err != nil {
		return GigSheet{}, err
	}

	return cols, nil
}

func GetSheetRows(path string) ([][]string, error) {
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

func getColumnsByType(rows [][]string) (GigSheet, error) {
	columnsByType := GigSheet{}
	if !checkSheetValidity(rows) {
		return GigSheet{}, fmt.Errorf("invalid sheet provided")
	}
	for _, row := range rows[1:] {
		columnsByType = appendValuesFromRowToColumns(columnsByType, row)
	}
	return columnsByType, nil
}

func checkSheetValidity(rows [][]string) bool {
	headingRow := rows[0]
	if headingRow[0] != "Bands" || headingRow[1] != "Venue" || headingRow[2] != "Date" ||
		headingRow[3] != "Who" || headingRow[4] != "Tour" || headingRow[5] != "Hotel" {
		return false
	}
	return true
}

func appendValuesFromRowToColumns(columnsByType GigSheet, row []string) GigSheet {
	if row[0] != "" {
		columnsByType.Bands = append(columnsByType.Bands, strings.Split(prettier.StandardiseCommaSeparated(row[0]), ", "))
	}

	if row[1] != "" {
		columnsByType.Venue = append(columnsByType.Venue, prettier.Standardise(row[1]))
	}

	if row[2] != "" {
		columnsByType.Date = append(columnsByType.Date, prettier.Standardise(row[2]))
	}

	if row[3] != "" {
		columnsByType.Who = append(columnsByType.Who, strings.Split(prettier.StandardiseCommaSeparated(row[3]), ", "))
	}

	if len(row) > 4 && row[4] != "" {
		columnsByType.Tour = append(columnsByType.Tour, prettier.Standardise(row[4]))
	}

	if len(row) > 5 && row[5] != "" {
		columnsByType.Hotel = append(columnsByType.Hotel, prettier.Standardise(row[5]))
	}

	return columnsByType
}

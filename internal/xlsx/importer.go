package xlsx

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func GetColumns(path string) (map[string][]string, error) {
	rows, err := getSheetRows(path)
	if err != nil {
		return map[string][]string{}, err
	}
	cols, err := getColumnsByType(rows)
	if err != nil {
		return map[string][]string{}, err
	}

	return cols, nil
}

func getSheetRows(path string) ([][]string, error) {
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

func getColumnsByType(rows [][]string) (map[string][]string, error) {
	columnsByType := map[string][]string{
		"Bands": {},
		"Venue": {},
		"Date":  {},
		"Who":   {},
		"Tour":  {},
		"Hotel": {},
	}
	if !checkSheetValidity(rows) {
		return map[string][]string{}, fmt.Errorf("invalid sheet provided")
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

func appendValuesFromRowToColumns(columnsByType map[string][]string, row []string) map[string][]string {
	if row[0] != "" {
		columnsByType["Bands"] = append(columnsByType["Bands"], row[0])
	}

	if row[1] != "" {
		columnsByType["Venue"] = append(columnsByType["Venue"], row[1])
	}

	if row[2] != "" {
		columnsByType["Date"] = append(columnsByType["Date"], row[2])
	}

	if row[3] != "" {
		columnsByType["Who"] = append(columnsByType["Who"], row[3])
	}

	if len(row) > 4 && row[4] != "" {
		columnsByType["Tour"] = append(columnsByType["Tour"], row[4])
	}

	if len(row) > 5 && row[5] != "" {
		columnsByType["Hotel"] = append(columnsByType["Hotel"], row[5])
	}

	return columnsByType
}

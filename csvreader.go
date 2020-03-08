// Package csvreader provides utility methods for easily reading a list of CSV files.
package csvreader

import (
	"encoding/csv"
	"os"
)

// ReadCSVFiles reads an array of filenames and combines them to return
// a two-dimensional array of results from the CSV.
func ReadCSVFiles(filenames []string, skipHeaders bool) ([][]string, error) {
	dataList := [][]string{}

	for _, filename := range filenames {
		csvLines, csvLinesErr := readCSVLines(filename)

		if csvLinesErr != nil {
			return [][]string{}, csvLinesErr
		}

		for i, csvLine := range csvLines {

			// Don't read first line
			if i == 0 && skipHeaders {
				continue
			}

			data := []string{}

			for _, csvColumn := range csvLine {
				data = append(data, csvColumn)
			}

			dataList = append(dataList, data)
		}
	}

	return dataList, nil
}

// readCSVlines is an internal function used for reading individual lines of a CSV file
func readCSVLines(filename string) ([][]string, error) {
	file, fileErr := os.Open(filename)

	if fileErr != nil {
		return [][]string{}, fileErr
	}

	defer file.Close()

	lines, linesErr := csv.NewReader(file).ReadAll()

	if linesErr != nil {
		return [][]string{}, linesErr
	}

	return lines, nil
}

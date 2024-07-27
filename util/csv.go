package util

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ProcessCSV(args []string, callback func(reader *csv.Reader) error) error {
		var file *os.File

		if IsPiped() {
			file = os.Stdin
		} else {
			var err error
			file, err = os.Open(args[0])
			if err != nil { return fmt.Errorf("Failed to open file: %s\n", err) }
			defer file.Close()
		}

		return callback(csv.NewReader(file))
}

func HashCSV(column string, path string) ([]uint32, error) {
	file, err := os.Open(path)
	if err != nil { return nil, fmt.Errorf("Failed to open file: %s\n", err) }
	defer file.Close()
	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil { return nil, err }

	col_index, err := ParseColumnFlag(column, headers)
	if err != nil { return nil, err }

	var hashes []uint32
	record := headers
	for {
		hashes = append(hashes, Hash(record[col_index]))

		record, err = reader.Read()
		if err != nil && err.Error() == "EOF" { break }
		if err != nil { return nil, err }
	}

	return hashes, nil
}

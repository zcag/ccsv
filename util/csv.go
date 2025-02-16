package util

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ProcessCSV(seperator rune, args []string, callback func(reader *csv.Reader) error) error {
		var file *os.File

		if IsPiped() {
			file = os.Stdin
		} else {
			var err error
			file, err = os.Open(args[0])
			if err != nil { return fmt.Errorf("Failed to open file: %s\n", err) }
			defer file.Close()
		}

		reader := csv.NewReader(file)
		reader.FieldsPerRecord = -1
		reader.Comma = seperator
		return callback(reader)
}

func HashCSV(separator rune, column string, path string) (map[uint32]struct{}, error) {
	file, err := os.Open(path)
	if err != nil { return nil, fmt.Errorf("Failed to open file: %s\n", err) }
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.Comma = separator

	headers, err := reader.Read()
	if err != nil { return nil, err }

	col_index, err := ParseColumnFlag(column, headers)
	if err != nil { return nil, err }

	hashes := make(map[uint32]struct{})
	record := headers
	for {
		hashes[Hash(record[col_index])] = struct{}{}
		record, err = reader.Read()
		if err != nil && err.Error() == "EOF" { break }
		if err != nil { return nil, err }
	}

	return hashes, nil
}

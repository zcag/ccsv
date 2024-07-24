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

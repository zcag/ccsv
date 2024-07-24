package util

import (
	"encoding/csv"
	"fmt"
	"strconv"
)

func ParseColumnFlag(
	reader *csv.Reader,
	columns []string,
) ([]int, error) {
	parsed := make([]int, len(columns))

	for i, col := range columns {
		num, err := strconv.Atoi(col)
		if err != nil { return nil, fmt.Errorf("Column by name is not yet supported.") }
		if num < 0 { return nil, fmt.Errorf("Column index cannot be smaller than zero")}
		parsed[i] = num
	}

	return parsed, nil
}

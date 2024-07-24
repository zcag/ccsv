package util

import (
	"fmt"
	"strconv"
)

func ParseColumnFlag(
	columns []string,
	headers []string,
) ([]int, error) {
	parsed := make([]int, len(columns))

	for i, col := range columns {
		num, err := strconv.Atoi(col)
		if err != nil {
			num, err = index(headers, col)
			if err != nil { return nil, fmt.Errorf("Can't find column %s.", col) }
		}

		if num < 0 { return nil, fmt.Errorf("Column index cannot be smaller than zero")}
		parsed[i] = num
	}

	return parsed, nil
}

func index(ar []string, item string) (int, error) {
	for i, cur := range ar {
		if cur == item { return i, nil }
	}

	return -1, fmt.Errorf("Item is not in the list")
}

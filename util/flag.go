package util

import (
	"fmt"
	"strconv"
)

func ParseColumnFlags(
	columns []string,
	headers []string,
) ([]int, error) {
	parsed := make([]int, len(columns))

	for i, col := range columns {
		num, err := ParseColumnFlag(col, headers)
		if err != nil { return nil, err }
		parsed[i] = num
	}

	return parsed, nil
}

func ParseColumnFlag(column string, headers []string) (int, error) {
	num, err := strconv.Atoi(column)
	if err != nil {
		num, err = index(headers, column)
		if err != nil { return 0, fmt.Errorf("Can't find column %s.", column) }
	}

	if num < 0 { return 0, fmt.Errorf("Column index cannot be smaller than zero")}

	return num, nil
}

func index(ar []string, item string) (int, error) {
	for i, cur := range ar {
		if cur == item { return i, nil }
	}

	return -1, fmt.Errorf("Item is not in the list")
}

package cmd

import (
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/cagdassalur/ccsv/util"
	"github.com/spf13/cobra"
)

var (
	no_headers bool
)

type stat struct {
	name string
	_type string
	count int
	nullCount int

	uniqCounts map[string]int

	longestChar int
	longestVal string

	_min int
	_max int
	sum int
}

func addData(s *stat, cell string) {
	if cell == "" {
		s.nullCount += 1
		return
	}
	if s.uniqCounts == nil { s.uniqCounts = make(map[string]int) }
	s.count += 1

	s.uniqCounts[cell]++

	if s._type == "number" || s._type == "" {
		num, err := strconv.Atoi(cell)
		if err == nil {
			s._type = "number"
			addNumberData(s, num)
			return
		}
	}

	s._type = "string"
	addStringData(s, cell)
}

func addStringData(s *stat, cell string) {
	n := len(cell)
	if n > s.longestChar {
		s.longestChar = n
		s.longestVal = cell
	}
}

func addNumberData(s *stat, cell int) {
	if cell < s._min { s._min = cell }
	if cell > s._max { s._max = cell }
	s.sum += cell
}


func printStat(_ int, s stat) {
	fmt.Println("-------------------------")
	fmt.Printf("Column: %s\n", s.name)
	fmt.Printf("Type: %s\n", s._type)
	fmt.Printf("Non-null values: %d\n", s.count)
	fmt.Printf("Null values: %d\n", s.nullCount)
	fmt.Printf("Uniq values: %d\n", len(s.uniqCounts))
	fmt.Printf("Repeating values: %d\n", s.count - len(s.uniqCounts))

	if s._type == "string" || len(s.uniqCounts) == s.count {
		fmt.Printf("Longest string length: %d\n", s.longestChar)
	} else if s._type == "number" {
		fmt.Printf("Min: %d\n", s._min)
		fmt.Printf("Max: %d\n", s._max)
		fmt.Printf("Mean: %2.f\n", float64(s.sum) / float64(s.count) )
		fmt.Printf("Sum: %d\n", s.sum)
	}

}

var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Show stats by column",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := util.ProcessCSV(args, func(reader *csv.Reader) error {

			var stats []stat 

			for row_i := 0; true; row_i++ {
				record, err := reader.Read()
				if err != nil && err.Error() == "EOF" { break }
				if err != nil { return err }

				for col_i, cell := range record {
					if row_i == 0 { 
						if !no_headers { 
							stats = append(stats, stat{name: cell})
							continue 
						}

						stats = append(stats, stat{name: fmt.Sprintf("%d", col_i)})
					}

					addData(&stats[col_i], cell)
				}
			}

			for i, s := range stats { printStat(i, s) }
			fmt.Println("")

			return nil
		})
		return err
	},
}

var headersCmd = &cobra.Command{
	Use:   "headers",
	Short: "Show headers and indexes",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := util.ProcessCSV(args, func(reader *csv.Reader) error {
			headers, err := reader.Read()
			if err != nil { return err }

			for i, header := range headers { fmt.Printf("%2d: %s\n", i, header) }

			return nil
		})
		return err
	},
}


func init() {
	rootCmd.AddCommand(statCmd)
	rootCmd.AddCommand(headersCmd)

	statCmd.Flags().BoolVarP(
		&no_headers,
		"no-headers",
		"H",
		false,
		"Do not parse first row as headers",
	)
}

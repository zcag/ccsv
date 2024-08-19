package cmd

import (
	"encoding/csv"
	"os"
	"regexp"

	"github.com/cagdassalur/ccsv/util"

	"github.com/spf13/cobra"
)

var (
	column_flag_match string
)

var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "Filter by matching regex on columns",
	Long: `Filter by matching regex on columns

	ccsv match -c name '\w+_\d' file.csv`,
	PreRunE: util.ValidateArgOrPipe("no input provided or piped; usage: ccsv cut -c[col,] [file]"),
	RunE: func(cmd *cobra.Command, args []string) error {
		pattern := args[0]
		err := util.ProcessCSV(args[1:], func(reader *csv.Reader) error {
			headers, err := reader.Read()
			if err != nil { return err }

			column, err := util.ParseColumnFlag(column_flag_match, headers)
			if err != nil { return err }

			writer := csv.NewWriter(os.Stdout)
			if err := writer.Write(headers); err != nil { return err }
			writer.Flush()

			for {
				record, err := reader.Read()
				if err != nil && err.Error() == "EOF" { break }
				if err != nil { return err }

				m, err := regexp.MatchString(pattern, record[column])
				if err != nil { return err }

				if m {
					if err := writer.Write(record); err != nil { return err }
					writer.Flush()
				}
			}

			return nil
		})

		return err
	},
}

func init() {
	rootCmd.AddCommand(matchCmd)

	matchCmd.Flags().StringVarP(
		&column_flag_match,
		"column",
		"c",
		"",
		"column name or index",
	)
	matchCmd.MarkFlagRequired("column")
}

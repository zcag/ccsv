package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"github.com/cagdassalur/ccsv/util"

	"github.com/spf13/cobra"
)

var (
	columns_flag []string
)

var cutCmd = &cobra.Command{
	Use: "cut -c [col] [file]",
	Short: "Cuts a csv by given columns by index or name",
	Long: `Cuts a csv by given columns by index or name
ccsv cut -c 1 some.csv
ccsv cut -c id some.csv
ccsv cut -c id -c 5 -c name some.csv`,
	Args: cobra.MaximumNArgs(1),
	PreRunE: util.ValidateArgOrPipe("no input provided or piped; usage: ccsv cut -c[col,] [file]"),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := processCSV(args, func(reader *csv.Reader) error {
			columns, err := util.ParseColumnFlag(reader, columns_flag)
			if err != nil { return err }

			writer := csv.NewWriter(os.Stdout)

			for {
				record, err := reader.Read()
				if err != nil && err.Error() == "EOF" { break }
				if err != nil { return err }

				outCells := make([]string, len(columns))
				for i, col := range columns {
					if col < len(record) { outCells[i] = record[col] } 
				}

				if err := writer.Write(outCells); err != nil { return err }
				writer.Flush()
			}

			return nil
		})

		return err
	},
}

func init() {
	rootCmd.AddCommand(cutCmd)

	cutCmd.Flags().StringArrayVarP(
		&columns_flag,
		"columns",
		"c",
		[]string{},
		"list of column names or indexes",
	)
	cutCmd.MarkFlagRequired("columns")
}

func processCSV(args []string, callback func(reader *csv.Reader) error) error {
		var reader *csv.Reader
		if util.IsPiped() {
			reader = csv.NewReader(os.Stdin)
		} else {
			file, err := os.Open(args[0])
			if err != nil { return fmt.Errorf("Failed to open file: %s\n", err) }
			defer file.Close()
			reader = csv.NewReader(file)
		}

		return callback(reader)
}

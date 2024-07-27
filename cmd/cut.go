package cmd

import (
	"encoding/csv"
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
ccsv cut -c id -c 5 -c age some.csv`,
	Args: cobra.MaximumNArgs(1),
	PreRunE: util.ValidateArgOrPipe("no input provided or piped; usage: ccsv cut -c[col,] [file]"),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := util.ProcessCSV(args, func(reader *csv.Reader) error {
			headers, err := reader.Read()
			if err != nil { return err }

			columns, err := util.ParseColumnFlags(columns_flag, headers)
			if err != nil { return err }

			writer := csv.NewWriter(os.Stdout)

			record := headers
			for {
				outCells := make([]string, len(columns))
				for i, col := range columns {
					if col < len(record) { outCells[i] = record[col] } 
				}

				if err := writer.Write(outCells); err != nil { return err }
				writer.Flush()

				record, err = reader.Read()
				if err != nil && err.Error() == "EOF" { break }
				if err != nil { return err }
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

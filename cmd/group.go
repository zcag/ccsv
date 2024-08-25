/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"encoding/csv"
	"regexp"

	"github.com/zcag/ccsv/util"

	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Output grouped csv files by column",
	Long: `Output grouped csv files by column

	ccsv group 'records_<country>' all_records.csv
	cat some.csv | ccsv group 'records_<3>'`,
	PreRunE: util.ValidateArgOrPipe("no input provided or piped; usage: ccsv cut -c[col,] [file]"),
	RunE: func(cmd *cobra.Command, args []string) error {
		output_pattern := args[0]
		err := util.ProcessCSV(args[1:], func(reader *csv.Reader) error {
			headers, err := reader.Read()
			if err != nil { return err }

			regex := regexp.MustCompile(`<(.*)>`)
			match := regex.FindStringSubmatch(output_pattern)
			if len(match) != 2 { 
				return fmt.Errorf("Provide a pattarn with column inside angle brackets. ex: output_<name>.csv") 
			}
			column, err := util.ParseColumnFlag(match[1], headers)
			if err != nil { return err }

			writers := make(map[string]*csv.Writer)
			filename_placeholder_regex := regexp.MustCompile(`(<.*>)`)
			for {
				record, err := reader.Read()
				if err != nil && err.Error() == "EOF" { break }
				if err != nil { return err }

				filename := filename_placeholder_regex.ReplaceAllString(output_pattern, record[column])
				writer, exists := writers[filename]
				if !exists {
					file, err := os.Create(filename)
					if err != nil { return err }
					writer = csv.NewWriter(file)

					if err := writer.Write(headers); err != nil { return err }
					writer.Flush()

					writers[filename] = writer
					defer file.Close()
				}

				if err := writer.Write(record); err != nil { return err }
				writer.Flush()
			}

			fmt.Printf("Done. Created %d files.. \n", len(writers))

			return nil
		})

		return err
	},
}

func init() {
	rootCmd.AddCommand(groupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

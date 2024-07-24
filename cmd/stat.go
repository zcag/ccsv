package cmd

import (
	"encoding/csv"
	"fmt"

	"github.com/cagdassalur/ccsv/util"
	"github.com/spf13/cobra"
)

var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Show stats by column",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stat called")
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
}

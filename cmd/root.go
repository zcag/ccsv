/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
    separator string
)

var rootCmd = &cobra.Command{
	Use: "ccsv",
	Short: "CLI tool for working with CSV files",
	Long: `
	ccsv [command] [opts] file.csv
	cat file.csv | ccsv [command] [opts]
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil { os.Exit(1) }
}

func init() {
	rootCmd.PersistentFlags().StringVar(&separator, "sep", ",", "CSV separator character")
}

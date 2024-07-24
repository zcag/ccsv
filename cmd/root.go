/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "ccsv",
	Short: "CLI tool for working with CSV files",
	Long: "",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil { os.Exit(1) }
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	"path/filepath"
	"bufio"
	"io"
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

const headerFile = "ccsv_header.tmp"

var headerSkipCmd = &cobra.Command{
	Use:   "header-skip",
	Aliases: []string{"hs"},
	Short: "Outputs the csv without the first header, saves the skipped line to be restored later",
	Long: `cat file.csv | ccsv header-skip | sort | ccsv header-restore`,
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		header, err := reader.ReadString('\n')
		if err != nil { return err }

		file := filepath.Join(os.TempDir(), headerFile)
		err = os.WriteFile(file, []byte(header), 0600)
		if err != nil { return err }


		_, err = io.Copy(os.Stdout, reader)
		return err
	},
}

var headerRestoreCmd = &cobra.Command{
	Use:   "header-restore",
	Aliases: []string{"hr"},
	Short: "Adds the headers previously saved to the input",
	Long: `cat file.csv | ccsv header-skip | sort | ccsv header-restore`,
	RunE: func(cmd *cobra.Command, args []string) error {
		file := filepath.Join(os.TempDir(), headerFile)
		headers, err := os.ReadFile(file)
		if err != nil { return err }

		os.Stdout.Write(headers)
		_, err = io.Copy(os.Stdout, os.Stdin)
		return err
	},
}

var headerAddCmd = &cobra.Command{
	Use:   "header-add",
	Aliases: []string{"ha"},
	Short: "Adds headers to a csv pipe",
	Long: `cat file.csv | ccsv ha id,name,age`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print(args[0])
		fmt.Print("\n")
		_, err := io.Copy(os.Stdout, os.Stdin)
		return err
	},
}

func init() {
	rootCmd.AddCommand(headerSkipCmd)
	rootCmd.AddCommand(headerRestoreCmd)
	rootCmd.AddCommand(headerAddCmd)
}

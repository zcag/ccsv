package util

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func ValidateArgOrPipe(errorMessage string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if	!IsPiped() && len(args) < 1 { return fmt.Errorf(errorMessage) }
		return nil
	}
}

func IsPiped() bool {
		stat, _ := os.Stdin.Stat()
		return (stat.Mode() & os.ModeCharDevice) == 0
}

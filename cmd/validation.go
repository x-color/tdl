package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func validateNumberArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}
	// Validate arguments are numbers
	for i, arg := range args {
		if _, err := strconv.Atoi(arg); err != nil {
			return fmt.Errorf("requires one or more numbers, argument %d is not number", i+1)
		}
	}
	return nil
}

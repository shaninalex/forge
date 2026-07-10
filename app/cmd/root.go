package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/forgecore/app/cmd/commands"
)

func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "forge",
	}

	// Commands
	cmd.AddCommand(commands.NewRootExecuteCommand())

	// Flags
	cmd.PersistentFlags().Bool("verbose", false, "Verbose output")

	return cmd
}

// Execute run application
func Execute() int {
	c := NewRootCmd()

	if err := c.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "forge",
	}

	// TODO:
	// Add commands here like this
	// cmd.AddCommand(commands.NewRootServeCommand())

	// NOTE:
	// - not sure about config path. For testing tool will be better have both - env and file based YAML
	// cmd.PersistentFlags().String("config", "", "Configuration path. Required.")

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

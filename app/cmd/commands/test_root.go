package commands

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

func NewRootExecuteCommand() (cmd *cobra.Command) {
	return &cobra.Command{
		Use:   "execute",
		Short: "execute given pipeline file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()

			_, err := os.Stat(args[0])
			if err != nil {
				panic(err)
			}

			ctx, appCancel := context.WithCancel(cmd.Context())
			defer appCancel()
			_ = c.Provide(func() context.Context {
				return ctx
			})

			err = c.Invoke(func(ctx context.Context) {

				log.Println("Executing pipeline")
			})
			if err != nil {
				panic(err)
			}
		},
	}
}

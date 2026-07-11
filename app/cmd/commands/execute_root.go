package commands

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/forgecore/app/executor"
	"go.uber.org/dig"
)

func NewRootExecuteCommand() (cmd *cobra.Command) {
	return &cobra.Command{
		Use:   "execute",
		Short: "execute given pipeline file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()

			pipelinePath := args[0]
			_, err := os.Stat(pipelinePath)
			if err != nil {
				panic(err)
			}

			ctx, appCancel := context.WithCancel(cmd.Context())
			defer appCancel()

			_ = c.Provide(func() context.Context {
				return ctx
			})

			_ = executor.Module(c)

			if err = c.Invoke(func(ctx context.Context, executor executor.Executor) {
				log.Println("Executing pipeline: ", pipelinePath)

				if err = executor.Parse(pipelinePath); err != nil {
					panic(err)
				}

				executor.Exec()

			}); err != nil {
				panic(err)
			}
		},
	}
}

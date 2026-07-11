package executor

import (
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(ProvideExecutor)

	return nil
}

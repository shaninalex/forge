package executor

import (
	"fmt"

	"gitlab.com/shaninalex/forgecore/app/model"
)

type Executor interface {
	Run(p *model.Pipeline)
	Exec(action *model.Action)
}

type BaseExecutor struct{}

var _ Executor = (*BaseExecutor)(nil)

func (s *BaseExecutor) Run(p *model.Pipeline) {
	for _, step := range p.Steps {
		s.Exec(&step.Action)
	}
}

func (s *BaseExecutor) Exec(action *model.Action) {
	fmt.Println("Executing action:", action)
}

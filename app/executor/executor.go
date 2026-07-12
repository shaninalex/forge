package executor

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gitlab.com/shaninalex/forgecore/app/executor/actions"
	"gitlab.com/shaninalex/forgecore/app/model"
	"gopkg.in/yaml.v3"
)

type Executor interface {
	Exec()
	Parse(pipeLinePath string) error
}

func ProvideExecutor() Executor {
	return &BaseExecutor{
		pipeline: nil,
		mux:      sync.Mutex{},
		results:  map[string]*model.Response{},
	}
}

type BaseExecutor struct {
	pipeline *model.Pipeline
	mux      sync.Mutex
	results  map[string]*model.Response
}

var _ Executor = (*BaseExecutor)(nil)

func (s *BaseExecutor) Exec() {
	t := time.Now()
	for _, step := range s.pipeline.Steps {
		switch a := step.Action.(type) {
		case *model.HttpAction:
			data, err := actions.ProcessHttpAction(a)
			if err != nil {
				fmt.Println("Error processing http action:", err)
				continue
			}

			// redirects are not implemented yet
			if data.StatusCode > 300 {
				fmt.Printf("[%s] Error processing http action: %d\nResponse: %s", step.Id, data.StatusCode, string(data.Body))
			}
			s.mux.Lock()
			s.results[step.Id] = data
			s.mux.Unlock()
		}

		if len(step.Asserts) > 0 {
			if !s.DoAssert(&step) {
				return
			}
		}
	}

	fmt.Printf("Executing action completed after: %fs\n", time.Since(t).Seconds())
}

func (s *BaseExecutor) DoAssert(step *model.Step) bool {
	fmt.Println("Checking step:", step.Id)
	fmt.Println("Executing step:", step.Asserts)
	return true
}

func (s *BaseExecutor) Parse(pipeLinePath string) error {
	var p model.Pipeline
	b, err := os.ReadFile(pipeLinePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &p)
	s.pipeline = &p
	return err
}

package executor

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
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
		st := time.Now()
		fmt.Printf("[%s]: executing step\n", step.Id)
		switch a := step.Action.(type) {
		case *model.HttpAction:
			data, err := ProcessHttpAction(a, s.results)
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
			if !s.DoAsserts(&step) {
				return
			}
		}
		fmt.Printf("[%s]: done after %fs\n", step.Id, time.Since(st).Seconds())
	}

	fmt.Printf("Executing action completed after: %fs\n", time.Since(t).Seconds())
}

func (s *BaseExecutor) DoAsserts(step *model.Step) bool {
	fmt.Println("Process asserts...")

	data, ok := s.results[step.Id]
	if !ok {
		return false
	}

	for i, a := range step.Asserts {
		pAssert, err := ParseAssert(a)
		if err != nil {
			fmt.Printf("[%d] Error parsing assert:\n%v\n", i, err)
			continue
		}
		res := ProcessAssert(data, pAssert)
		if !res {
			fmt.Print(color.RedString("[%d] %s: %t\n", i, a, res))
			return false
		}

		fmt.Print(color.GreenString("[%d] %s: %t\n", i, a, res))
	}

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

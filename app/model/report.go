package model

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Pipeline struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Id      string   `yaml:"id"`
	Action  Action   `yaml:"action"`
	Asserts []string `yaml:"assert"`
}

func (s *Step) UnmarshalYAML(node *yaml.Node) error {
	var probe struct {
		Id      string     `yaml:"id"`
		Type    ActionType `yaml:"type"`
		Asserts []string   `yaml:"assert"`
	}

	if err := node.Decode(&probe); err != nil {
		return err
	}

	s.Id = probe.Id
	s.Asserts = probe.Asserts

	switch probe.Type {
	case ActionHTTP:
		var a HttpAction
		if err := node.Decode(&a); err != nil {
			return err
		}
		s.Action = &a

	case ActionGraphQL:
		var a GraphQLAction
		if err := node.Decode(&a); err != nil {
			return err
		}
		s.Action = &a

	case ActionWebhookListener:
		var a WebhookListenerAction
		if err := node.Decode(&a); err != nil {
			return err
		}
		s.Action = &a

	case ActionMcpCall:
		var a McpCallAction
		if err := node.Decode(&a); err != nil {
			return err
		}
		s.Action = &a

	default:
		return fmt.Errorf("unknown action type %q", probe.Type)
	}

	return nil
}

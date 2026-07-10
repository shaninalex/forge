package model

type Pipeline struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Id      string   `yaml:"id"`
	Action  Action   `yaml:"action"`
	Asserts []Assert `yaml:"asserts"`
}

package executor_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/forgecore/app/executor"
)

func Test_Tokenizer(t *testing.T) {
	var expressions = []string{
		"headers X-User eq 1",
		"body id eq 1",
		"status lt 400",
		"body status eq true",
	}

	for _, expression := range expressions {
		a, err := executor.ParseAssert(expression)
		assert.Nil(t, err)
		assert.NotNil(t, a)
	}
}

func Test_FindParams(t *testing.T) {
	expressions := []struct {
		exp    string
		result []string
	}{
		{
			exp: `{
		        "title": "foo",
		        "body": "bar",
		        "userId": {{ user id }}
			}`,
			result: []string{"{{ user id }}"},
		},
		{
			exp: `{
		        "title": "foo",
		        "email": {{ user email }},
		        "userId": {{ user id }}
			}`,
			result: []string{"{{ user email }}", "{{ user id }}"},
		},
		{
			exp:    "",
			result: nil,
		},
		{
			exp:    "hello world",
			result: nil,
		},
	}

	for _, expression := range expressions {
		res := executor.FindParams(expression.exp)
		assert.True(t, reflect.DeepEqual(res, expression.result))
	}
}

var testParamsData map[string]string = map[string]string{
	"step_a": `{
		"value_a": "value a from step a",
		"value_b": "value b from step a",
	}`,
	"step_b": `{
		"value_c": "value c from step b",
		"value_d": "value d from step b",
	}`,
	"step_c": `{
		"value_e": "value e from step c",
		"value_f": {
			"value_g": "value g from step c"
		},
	}`,
}

func Test_ParseParams(t *testing.T) {
	expressions := []struct {
		exp    string
		result string
	}{
		{
			exp:    `{{ step_b | value_c }}`,
			result: "value c from step b",
		},
		{
			exp:    `{{ step_c | value_f.value_g }}`,
			result: "value g from step c",
		},
	}

	for _, expression := range expressions {
		res := executor.ProcessParams(expression.exp, testParamsData)
		assert.True(t, reflect.DeepEqual(res, expression.result))
	}
}

func Test_ParseParams_PanicsNoData(t *testing.T) {
	assert.PanicsWithError(t,
		executor.CantFindDataProcessParamsError.Error(),
		func() {
			_ = executor.ProcessParams(`{{ not_exist | a }}`, testParamsData)
		},
	)
}

func Test_ParseParams_PanicsNoValue(t *testing.T) {
	assert.PanicsWithError(t,
		executor.CantFindValueProcessParamsError.Error(),
		func() {
			_ = executor.ProcessParams(`{{ step_c | no_value }}`, testParamsData)
		},
	)
}

func Test_ApplyParams(t *testing.T) {
	expressions := []struct {
		exp    string
		result string
	}{
		{
			exp:    `hello {{ step_b | value_c }}`,
			result: "hello value c from step b",
		},
		{
			exp:    `hello {{ step_c | value_f.value_g }}`,
			result: "hello value g from step c",
		},
	}

	for _, expression := range expressions {
		res := executor.ApplyParams(expression.exp, testParamsData)
		assert.True(t, reflect.DeepEqual(res, expression.result))
	}
}

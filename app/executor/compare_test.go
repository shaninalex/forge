package executor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/forgecore/app/executor"
	"gitlab.com/shaninalex/forgecore/app/model"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		name     string
		actual   any
		expected any
		op       model.CompareOperator
		want     bool
	}{
		// Equality
		{"eq int", 1, 1, model.IsEqualTo, true},
		{"eq string", "abc", "abc", model.IsEqualTo, true},
		{"eq bool", true, true, model.IsEqualTo, true},
		{"eq slice", []int{1, 2}, []int{1, 2}, model.IsEqualTo, true},
		{"eq map", map[string]int{"a": 1}, map[string]int{"a": 1}, model.IsEqualTo, true},
		{"eq different", 1, 2, model.IsEqualTo, false},
		{"eq int vs numeric string", 201, "201", model.IsEqualTo, true},
		{"eq float vs numeric string", float64(200), "200", model.IsEqualTo, true},
		{"eq int vs non-matching string", 201, "200", model.IsEqualTo, false},
		{"eq bool vs string", true, "true", model.IsEqualTo, true},

		// Not equal
		{"neq", 1, 2, model.IsNotEqualTo, true},
		{"neq false", "abc", "abc", model.IsNotEqualTo, false},
		{"neq int vs numeric string false", 201, "201", model.IsNotEqualTo, false},

		// Greater
		{"gt int", 10, 5, model.IsGreater, true},
		{"gt float", 10.5, 10.1, model.IsGreater, true},
		{"gt false", 5, 10, model.IsGreater, false},
		{"gt incompatible", "10", 5, model.IsGreater, false},

		// Greater or equal
		{"gte equal", 10, 10, model.IsGreaterOrEqual, true},
		{"gte greater", 11, 10, model.IsGreaterOrEqual, true},
		{"gte false", 9, 10, model.IsGreaterOrEqual, false},

		// Lower
		{"lt", 5, 10, model.IsLower, true},
		{"lt false", 10, 5, model.IsLower, false},

		// Lower or equal
		{"lte equal", 10, 10, model.IsLowerOrEqual, true},
		{"lte lower", 5, 10, model.IsLowerOrEqual, true},
		{"lte false", 11, 10, model.IsLowerOrEqual, false},

		// Contains
		{"contains", "hello world", "world", model.IsContain, true},
		{"contains false", "hello", "world", model.IsContain, false},
		{"contains wrong type", 123, "23", model.IsContain, false},

		// Not contains
		{"ncontains", "hello", "world", model.IsNotContain, true},
		{"ncontains false", "hello", "ell", model.IsNotContain, false},

		// Starts
		{"starts", "golang", "go", model.IsStartWith, true},
		{"starts false", "golang", "lang", model.IsStartWith, false},

		// Not starts
		{"nstarts", "golang", "lang", model.IsNotStartWith, true},
		{"nstarts false", "golang", "go", model.IsNotStartWith, false},

		// Ends
		{"ends", "golang", "lang", model.IsEndWith, true},
		{"ends false", "golang", "go", model.IsEndWith, false},

		// Not ends
		{"nends", "golang", "go", model.IsNotEndWith, true},
		{"nends false", "golang", "lang", model.IsNotEndWith, false},

		// Empty
		{"empty nil", nil, nil, model.IsEmpty, true},
		{"empty string", "", nil, model.IsEmpty, true},
		{"empty slice", []int{}, nil, model.IsEmpty, true},
		{"empty map", map[string]int{}, nil, model.IsEmpty, true},
		{"empty false", "abc", nil, model.IsEmpty, false},

		// Not empty
		{"nempty string", "abc", nil, model.IsNotEmpty, true},
		{"nempty slice", []int{1}, nil, model.IsNotEmpty, true},
		{"nempty false", "", nil, model.IsNotEmpty, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := executor.Compare(tt.actual, tt.expected, tt.op)
			assert.Equal(t, tt.want, got)
		})
	}
}

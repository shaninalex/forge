package model

type FilterCondition string

const (
	IsEqualTo        FilterCondition = "eq"
	IsNotEqualTo     FilterCondition = "neq"
	IsGreater        FilterCondition = "gt"
	IsGreaterOrEqual FilterCondition = "gteq"
	IsLower          FilterCondition = "lt"
	IsLowerOrEqual   FilterCondition = "lteq"
	IsContain        FilterCondition = "contains"
	IsNotContain     FilterCondition = "ncontains"
	IsStartWith      FilterCondition = "starts"
	IsNotStartWith   FilterCondition = "nstarts"
	IsEndWith        FilterCondition = "ends"
	IsNotEndWith     FilterCondition = "nends"
	IsEmpty          FilterCondition = "empty"
	IsNotEmpty       FilterCondition = "nempty"
	IsTimeInterval   FilterCondition = "time_interval"
	IsTimePart       FilterCondition = "time_part"
)

var TFilterCondition = []FilterCondition{
	IsEqualTo,
	IsNotEqualTo,
	IsGreater,
	IsGreaterOrEqual,
	IsLower,
	IsLowerOrEqual,
	IsContain,
	IsNotContain,
	IsStartWith,
	IsNotStartWith,
	IsEndWith,
	IsNotEndWith,
	IsEmpty,
	IsNotEmpty,
	IsTimeInterval,
	IsTimePart,
}

func (c FilterCondition) IsValid() bool {
	for _, v := range TFilterCondition {
		if v == c {
			return true
		}
	}
	return false
}

// Assert the core of "tokenizer".
type Assert struct {
	Expression string

	Op    FilterCondition
	Query string
	Value string
}

type Query struct {
	Expression string
}

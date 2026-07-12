package model

type CompareOperator string

const (
	IsEqualTo        CompareOperator = "eq"
	IsNotEqualTo     CompareOperator = "neq"
	IsGreater        CompareOperator = "gt"
	IsGreaterOrEqual CompareOperator = "gteq"
	IsLower          CompareOperator = "lt"
	IsLowerOrEqual   CompareOperator = "lteq"
	IsContain        CompareOperator = "contains"
	IsNotContain     CompareOperator = "ncontains"
	IsStartWith      CompareOperator = "starts"
	IsNotStartWith   CompareOperator = "nstarts"
	IsEndWith        CompareOperator = "ends"
	IsNotEndWith     CompareOperator = "nends"
	IsEmpty          CompareOperator = "empty"
	IsNotEmpty       CompareOperator = "nempty"
	IsTimeInterval   CompareOperator = "time_interval"
	IsTimePart       CompareOperator = "time_part"
)

var TCompareOperator = []CompareOperator{
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

func (c CompareOperator) IsValid() bool {
	for _, v := range TCompareOperator {
		if v == c {
			return true
		}
	}
	return false
}

// Assert the core of "tokenizer".
type Assert struct {
	Expression string

	Op    CompareOperator
	Query string
	Value string
}

type Query struct {
	Expression string
}

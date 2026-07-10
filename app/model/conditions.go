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

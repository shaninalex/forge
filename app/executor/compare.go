package executor

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"gitlab.com/shaninalex/forgecore/app/model"
)

func Compare(actual, expected any, op model.CompareOperator) bool {
	switch op {
	case model.IsEqualTo:
		return looseEqual(actual, expected)

	case model.IsNotEqualTo:
		return !looseEqual(actual, expected)

	case model.IsEmpty:
		return isEmpty(actual)

	case model.IsNotEmpty:
		return !isEmpty(actual)

	case model.IsContain:
		a, ok1 := actual.(string)
		e, ok2 := expected.(string)
		return ok1 && ok2 && strings.Contains(a, e)

	case model.IsNotContain:
		a, ok1 := actual.(string)
		e, ok2 := expected.(string)
		return ok1 && ok2 && !strings.Contains(a, e)

	case model.IsStartWith:
		a, ok1 := actual.(string)
		e, ok2 := expected.(string)
		return ok1 && ok2 && strings.HasPrefix(a, e)

	case model.IsNotStartWith:
		a, ok1 := actual.(string)
		e, ok2 := expected.(string)
		return ok1 && ok2 && !strings.HasPrefix(a, e)

	case model.IsEndWith:
		a, ok1 := actual.(string)
		e, ok2 := expected.(string)
		return ok1 && ok2 && strings.HasSuffix(a, e)

	case model.IsNotEndWith:
		a, ok1 := actual.(string)
		e, ok2 := expected.(string)
		return ok1 && ok2 && !strings.HasSuffix(a, e)

	case model.IsGreater:
		return compareNumbers(actual, expected, func(a, b float64) bool { return a > b })

	case model.IsGreaterOrEqual:
		return compareNumbers(actual, expected, func(a, b float64) bool { return a >= b })

	case model.IsLower:
		return compareNumbers(actual, expected, func(a, b float64) bool { return a < b })

	case model.IsLowerOrEqual:
		return compareNumbers(actual, expected, func(a, b float64) bool { return a <= b })

	case model.IsTimeInterval:
		panic("not implemented")

	case model.IsTimePart:
		panic("not implemented")

	default:
		return false
	}
}

func looseEqual(actual, expected any) bool {
	if reflect.DeepEqual(actual, expected) {
		return true
	}

	if a, ok := asNumber(actual); ok {
		if b, ok := asNumber(expected); ok {
			return a == b
		}
	}

	return fmt.Sprint(actual) == fmt.Sprint(expected)
}

func asNumber(v any) (float64, bool) {
	if s, ok := v.(string); ok {
		f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		return f, err == nil
	}
	return toFloat64(v)
}

func compareNumbers(actual, expected any, cmp func(float64, float64) bool) bool {
	a, ok := toFloat64(actual)
	if !ok {
		return false
	}

	b, ok := toFloat64(expected)
	if !ok {
		return false
	}

	return cmp(a, b)
}

func toFloat64(v any) (float64, bool) {
	switch n := v.(type) {
	case int:
		return float64(n), true
	case int8:
		return float64(n), true
	case int16:
		return float64(n), true
	case int32:
		return float64(n), true
	case int64:
		return float64(n), true
	case uint:
		return float64(n), true
	case uint8:
		return float64(n), true
	case uint16:
		return float64(n), true
	case uint32:
		return float64(n), true
	case uint64:
		return float64(n), true
	case float32:
		return float64(n), true
	case float64:
		return n, true
	default:
		return 0, false
	}
}

func isEmpty(v any) bool {
	if v == nil {
		return true
	}

	switch x := v.(type) {
	case string:
		return x == ""
	case bool:
		return false
	}

	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		return rv.Len() == 0
	case reflect.Pointer, reflect.Interface:
		return rv.IsNil()
	}

	return reflect.DeepEqual(v, reflect.Zero(rv.Type()).Interface())
}

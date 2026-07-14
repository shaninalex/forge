/*
This package parse assert expressions and convert's them into processable structs.
Examples of expressions:
- headers X-User eq 1
- body id eq 1
- status lt 400
- body status eq true
*/

package executor

import (
	"errors"
	"regexp"
	"strings"

	"github.com/tidwall/gjson"
	"gitlab.com/shaninalex/forgecore/app/model"
)

func ParseAssert(expression string) (*model.Assert, error) {
	expression = strings.TrimSpace(expression)
	parts := strings.Split(expression, " ")
	a := &model.Assert{
		Expression: expression,
	}
	for i, p := range parts {
		if !model.CompareOperator(p).IsValid() {
			continue
		}

		a.Op = model.CompareOperator(p)
		a.Query = strings.Join(parts[:i], " ")
		a.Value = strings.Join(parts[i+1:], " ")
	}
	return a, nil
}

const (
	paramsRegexp = `\{\{[^{}]*\}\}`
)

// FindParams search {{ ... }} in text and return a list of them
func FindParams(s string) []string {
	r, _ := regexp.Compile(paramsRegexp)
	return r.FindAllString(s, -1)
}

var (
	TooManyPartsProcessParamsError  = errors.New("too many parts")
	CantFindDataProcessParamsError  = errors.New("can't find data")
	CantFindValueProcessParamsError = errors.New("can't find param")
)

// ProcessParams parse params
func ProcessParams(s string, dataBank map[string]string) string {
	s = strings.Trim(s, " {{}}")
	parts := strings.Split(s, " | ")
	if len(parts) > 2 {
		panic(TooManyPartsProcessParamsError)
	}

	data, ok := dataBank[parts[0]]
	if !ok {
		panic(CantFindDataProcessParamsError)
	}

	result := gjson.Get(data, parts[1])

	if !result.Exists() {
		panic(CantFindValueProcessParamsError)
	}

	return result.String()
}

// ApplyParams find and replace placeholder {{...}} with values from dataBank
func ApplyParams(s string, dataBank map[string]string) string {
	params := FindParams(s)
	for _, p := range params {
		value := ProcessParams(p, dataBank)
		s = strings.ReplaceAll(s, p, value)
	}
	return s
}

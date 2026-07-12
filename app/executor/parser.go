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
	"strings"

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

package tokenizer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/forgecore/app/executor/tokenizer"
)

func Test_Tokenizer_Parse(t *testing.T) {
	var expressions = []string{
		"headers X-User eq 1",
		"body id eq 1",
		"status lt 400",
		"body status eq true",
	}

	for _, expression := range expressions {
		a, err := tokenizer.Parse(expression)
		assert.Nil(t, err)
		assert.NotNil(t, a)
	}
}

package executor

import (
	"errors"
	"log"
	"slices"
	"strings"

	"github.com/tidwall/gjson"
	"gitlab.com/shaninalex/forgecore/app/model"
)

func ProcessAssert(resp model.Response, a model.Assert) bool {
	// get data from response via a.Query
	_, err := GetData(resp, a.Query)
	if err != nil {
		log.Println(err)
		return false
	}

	// TODO: compare (a.Op) that data with a.Value

	return true
}

var (
	InvalidQueryError              = errors.New("invalid query")
	CantFindJsonKeyError           = errors.New("can't find json key")
	ContentTypeNotProcessableError = errors.New("content type is not processable")
	UnknownQueryPartsError         = errors.New("unknown query parts")
	CantGetDataError               = errors.New("can't get data")
)

var qParts = []string{
	"header", "body",
}

func GetData(resp model.Response, expression string) (any, error) {
	expression = strings.TrimSpace(expression)
	parts := strings.Split(expression, " ")
	l := len(parts)

	if l < 1 {
		return nil, InvalidQueryError
	}

	if l == 1 && parts[0] == "status" {
		return resp.StatusCode, nil
	}

	if !slices.Contains(qParts, parts[0]) {
		return nil, UnknownQueryPartsError
	}

	if parts[0] == "header" {
		return resp.Header.Get(parts[1]), nil
	}

	if parts[0] == "body" {
		if resp.Header.Get("Content-Type") != "application/json" {
			return nil, ContentTypeNotProcessableError
		}

		result := gjson.Get(string(resp.Body), parts[1])
		if !result.Exists() {
			return nil, CantFindJsonKeyError
		}

		return result.Value(), nil
	}

	return nil, CantGetDataError
}

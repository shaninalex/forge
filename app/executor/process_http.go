package executor

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.com/shaninalex/forgecore/app/model"
)

func ProcessHttpAction(action *model.HttpAction, data model.DataBank) (*model.Response, error) {
	var body io.Reader
	if action.Method == model.MethodPost ||
		action.Method == model.MethodPut ||
		action.Method == model.MethodPatch {
		body = strings.NewReader(action.Body)
	}

	u, err := url.Parse(action.Url)
	if err != nil {
		return nil, err
	}

	if action.Query != nil {
		q := u.Query()
		for k, v := range action.Query {
			params := FindParams(v)
			if len(params) > 0 {
				for _, p := range params {
					v = ApplyParams(p, data.GetDataMap())
				}
			} else {
				q.Set(k, v)
			}
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(string(action.Method), u.String(), body)
	if err != nil {
		return nil, err
	}

	if action.ContentType == "" {
		req.Header.Add("Content-Type", "application/json")
	} else {
		req.Header.Add("Content-Type", action.ContentType)
	}

	if action.Headers != nil {
		for k, v := range action.Headers {
			req.Header.Add(k, v)
		}
	}

	t := time.Now()
	resp, err := http.DefaultClient.Do(req)
	since := time.Since(t)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	return &model.Response{
		StatusCode: resp.StatusCode,
		Duration:   since.Milliseconds(),
		Header:     resp.Header,
		Body:       bodyBytes,
	}, nil
}

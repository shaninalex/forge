package executor_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/forgecore/app/executor"
	"gitlab.com/shaninalex/forgecore/app/model"
)

func Test_GetData(t *testing.T) {
	jsonResp := &model.Response{
		Body:       []byte(`{"id": 1, "name": "forge", "active": true, "nested": {"key": "value"}}`),
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
			"X-User":       []string{"25"},
		},
	}

	tests := []struct {
		name       string
		resp       *model.Response
		expression string
		want       any
		wantErr    error
	}{
		{
			name:       "status returns status code",
			resp:       jsonResp,
			expression: "status",
			want:       200,
		},
		{
			name:       "status trims surrounding whitespace",
			resp:       jsonResp,
			expression: "  status  ",
			want:       200,
		},
		{
			name:       "header returns header value",
			resp:       jsonResp,
			expression: "header X-User",
			want:       "25",
		},
		{
			name:       "header returns empty string for missing header",
			resp:       jsonResp,
			expression: "header X-Missing",
			want:       "",
		},
		{
			name:       "body returns numeric json value",
			resp:       jsonResp,
			expression: "body id",
			want:       float64(1),
		},
		{
			name:       "body returns string json value",
			resp:       jsonResp,
			expression: "body name",
			want:       "forge",
		},
		{
			name:       "body returns bool json value",
			resp:       jsonResp,
			expression: "body active",
			want:       true,
		},
		{
			name:       "body returns nested json value",
			resp:       jsonResp,
			expression: "body nested.key",
			want:       "value",
		},
		{
			name:       "body with missing json key errors",
			resp:       jsonResp,
			expression: "body missing",
			wantErr:    executor.CantFindJsonKeyError,
		},
		{
			name: "body with non-json content type errors",
			resp: &model.Response{
				Body:       []byte("plain text"),
				StatusCode: 200,
				Header: http.Header{
					"Content-Type": []string{"text/plain"},
				},
			},
			expression: "body id",
			wantErr:    executor.ContentTypeNotProcessableError,
		},
		{
			name:       "unknown query part errors",
			resp:       jsonResp,
			expression: "cookie session",
			wantErr:    executor.UnknownQueryPartsError,
		},
		{
			name:       "empty expression errors",
			resp:       jsonResp,
			expression: "",
			wantErr:    executor.UnknownQueryPartsError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executor.GetData(tt.resp, tt.expression)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, got)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

package tests

import (
	"bytes"
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/antonT001/easy-storage-light/internal/rest"
	"github.com/antonT001/easy-storage-light/internal/rest/tests/mocks"
	"github.com/stretchr/testify/assert"
)

const emptyBody string = ""

type (
	body        []byte
	requestBody interface {
		Reader() io.Reader
	}

	assertFn func(t *testing.T, resp *httptest.ResponseRecorder)
	mockFn   func(m *mocks.Mocks)

	testCase struct {
		name    string
		method  string
		path    string
		headers map[string]string
		body    requestBody
		prepare mockFn
		assert  assertFn
	}
)

func (b body) Reader() io.Reader {
	return bytes.NewReader(b)
}

func setBoby(bodyByte []byte) requestBody {
	return body(bodyByte)
}

func run(t *testing.T, tests []testCase) {
	for _, tt := range tests {
		api := setup(t, tt.prepare)

		t.Run(
			tt.name, func(t *testing.T) {
				resp := do(api, tt.method, tt.path, tt.headers, tt.body)
				tt.assert(t, resp)
			},
		)
	}
}

func do(h *rest.Server, method, path string, headers map[string]string, body requestBody) *httptest.ResponseRecorder {
	var reader io.Reader
	if body != nil {
		reader = body.Reader()
	}

	req := httptest.NewRequest(method, path, reader)

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	rr := httptest.NewRecorder()
	h.App.Handler.ServeHTTP(rr, req)

	return rr
}

func getBody(resp *httptest.ResponseRecorder) string {
	data, _ := io.ReadAll(resp.Body)
	return string(data)
}

func readFile(file string) string {
	data, _ := os.ReadFile(file)
	return string(data)
}

func assertResponse(code int, body string) assertFn {
	return func(t *testing.T, resp *httptest.ResponseRecorder) {
		assert.NotEmpty(t, resp)

		if !assert.Equal(t, code, resp.Code) {
			t.Log(getBody(resp))
		}

		if body != emptyBody {
			assert.JSONEq(t, body, getBody(resp))
		}
	}
}

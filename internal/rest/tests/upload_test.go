package tests

import (
	"net/http"
	"testing"

	"github.com/antonT001/easy-storage-light/internal/lib/httplib"
	"github.com/antonT001/easy-storage-light/internal/rest/tests/mocks"
)

func TestUpload(t *testing.T) {
	tests := []testCase{
		{
			name:   "valid request",
			method: http.MethodPost,
			path:   "/api/v1/files/upload",
			headers: map[string]string{
				httplib.UUIDHeaderKey:                "39c4739c-91cf-11ee-b9d1-0242ac120002",
				httplib.ChunkNumHeaderKey:            "0",
				httplib.NameHeaderKey:                "gopher.jpg",
				httplib.SHA256ChunkChecksumHeaderKey: "cbb756eb255316279a3e09cb7342c38754060a5b4bd6560e14f51d85cbd745e6",
				httplib.SHA256FileChecksumHeaderKey:  "cd372fb85148700fa88095e3492d3f9f5beb43e555e5ff26d95f5a6adc36f8e6",
			},
			body:    setBoby([]byte("not empty request body")),
			prepare: mocks.Upload,
			assert:  assertResponse(http.StatusOK, emptyBody),
		},
		{
			name:    "invalid headers",
			method:  http.MethodPost,
			path:    "/api/v1/files/upload",
			headers: map[string]string{}, // empty headers
			assert:  assertResponse(http.StatusBadRequest, readFile("./data/invalid-headers.json")),
		},
		{
			name:   "empty request body",
			method: http.MethodPost,
			path:   "/api/v1/files/upload",
			headers: map[string]string{
				httplib.UUIDHeaderKey:                "39c4739c-91cf-11ee-b9d1-0242ac120002",
				httplib.ChunkNumHeaderKey:            "1",
				httplib.NameHeaderKey:                "gopher.jpg",
				httplib.SHA256ChunkChecksumHeaderKey: "cbb756eb255316279a3e09cb7342c38754060a5b4bd6560e14f51d85cbd745e6",
				httplib.SHA256FileChecksumHeaderKey:  "cd372fb85148700fa88095e3492d3f9f5beb43e555e5ff26d95f5a6adc36f8e6",
			},
			assert: assertResponse(http.StatusBadRequest, readFile("./data/invalid-body.json")),
		},
		{
			name:    "method not allowed",
			method:  http.MethodGet,
			path:    "/api/v1/files/upload",
			headers: map[string]string{}, // empty headers
			assert:  assertResponse(http.StatusMethodNotAllowed, ""),
		},
	}
	run(t, tests)
}

package restClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func getMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

// Mock
func StartMockups() {
	enabledMocks = true
}

func FlushMockups() {
	mocks = make(map[string]*Mock)
}

func StopMockups() {
	enabledMocks = false
}

func AddMockup(mock Mock) {
	mocks[getMockId(mock.HttpMethod, mock.Url)] = &mock
}

// Basic Post api
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	// Mocking area ===================================================
	if enabledMocks {
		mock := mocks[getMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("No mockup found for give request")
		}
		return mock.Response, mock.Err
	}
	// ================================================================

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}

	return client.Do(request)
}

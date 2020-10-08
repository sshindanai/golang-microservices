package githib_provider

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/sshindanai/golang-microservices/src/api/clients/restClient"
	"github.com/sshindanai/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restClient.StartMockups()
	os.Exit(m.Run())
}

func TestConstant(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoRestClient(t *testing.T) {
	restClinet.FlushMocks()
	restClient.AddMockup(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("Invalid resclient response"),
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid resclient response", err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restClinet.FlushMocks()

	invalidCloser, _ := os.Open("-asf3")
	restClient.AddMockup(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusCreated, err.StatusCode)
	assert.EqualValues(t, "Invalid Response Body", err.Message)
}

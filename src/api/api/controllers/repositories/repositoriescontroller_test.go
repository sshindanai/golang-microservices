package repositories

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/sshindanai/golang-microservices/src/api/api/errors"
	"github.com/sshindanai/golang-microservices/src/api/clients/restClient"
	"github.com/sshindanai/golang-microservices/src/api/domain/repositories"
	"github.com/sshindanai/golang-microservices/src/api/utils/testutils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Mocking for real api call
	restClient.StartMockups()
	os.Exit(m.Run())
}

func TestInvalidJsonRequest(t *testing.T) {
	// Mocking gin-gonic
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := testutils.GetMockedContext(request, response)

	// Action- Call api
	CreateRepo(c)

	// Assertions
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "Invalid Json body", apiErr.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restClient.FlushMockups()
	restClient.AddMockup(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/docs"}`)),
		},
	})

	// Mocking gin-gonic
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c := testutils.GetMockedContext(request, response)

	// Action- Call api
	CreateRepo(c)

	// Assertions
	assert.EqualValues(t, http.StatusUnauthorized, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "Requires authentication", apiErr.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restClient.FlushMockups()
	restClient.AddMockup(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c := testutils.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}

package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/sshindanai/golang-microservices/src/api/clients/restClient"
	"github.com/sshindanai/golang-microservices/src/api/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restClient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "Invalid repository name", err.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restClient.FlushMockups()
	restClient.AddMockup(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
}

func TestCreateRepoNoError(t *testing.T) {
	restClient.FlushMockups()
	restClient.AddMockup(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name":"testing", "owner": {"login": "sshindanai"}}`)),
		},
	})

	//ioutil.NopCloser(strings.NewReader(`{"message: "Requires authentication", "documentation_url": "https://developer.github.com/docs"}`)),

	request := repositories.CreateRepoRequest{Name: "testing"}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "sshindanai", result.Owner)
}

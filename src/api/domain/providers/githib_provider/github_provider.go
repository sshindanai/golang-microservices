package githib_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sshindanai/golang-microservices/src/api/clients/restClient"
	"github.com/sshindanai/golang-microservices/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	//accessToken               = "016cd1906270c87fc65da8c1724eb64ee9e45c35"
	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := restClient.Post(urlCreateRepo, request, headers)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Response Body"}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid json response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "error when trying to unmarshal Github create repo response"}
	}

	return &result, nil
}

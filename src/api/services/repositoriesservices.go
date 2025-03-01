package services

import (
	"net/http"
	"sync"

	"github.com/sshindanai/golang-microservices/src/api/api/errors"
	"github.com/sshindanai/golang-microservices/src/api/config"
	"github.com/sshindanai/golang-microservices/src/api/domain/github"
	"github.com/sshindanai/golang-microservices/src/api/domain/providers/githib_provider"
	"github.com/sshindanai/golang-microservices/src/api/domain/repositories"
	"github.com/sshindanai/golang-microservices/src/log/optionb"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(clientId string, req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(clientId string, req []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

// For mocking purpose
func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	req := github.CreateRepoRequest{
		Name:        input.Name,
		Private:     false,
		Description: input.Description,
	}
	optionb.Info("about to send request to external api", optionb.Field("client_id", clientId),
		optionb.Field("status", "pending"))

	// Small bug on GetAccessToken
	response, err := githib_provider.CreateRepo(config.GithubAccessToken, req)
	if err != nil {
		optionb.Error("response obtain from external api", err, optionb.Field("client_id", clientId), optionb.Field("status", "error"))
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	optionb.Info("response obtain from external api", optionb.Field("client_id", clientId), optionb.Field("status", "success"))
	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}

func (s *repoService) CreateRepos(clientId string, req []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	// n requests to process
	for _, current := range req {
		wg.Add(1)
		go s.createRepoConcurrent(clientId, current, input)
	}

	// Wait until had all the results
	wg.Wait()
	close(input)

	result := <-output

	// Count for compairing the success results and the number of requests
	successCreation := 0
	for _, current := range result.Results {
		if current.Response != nil {
			// Success
			successCreation++
		}
	}
	if successCreation == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreation == len(req) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	// Results to send to channel
	var results repositories.CreateReposResponse

	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *repoService) createRepoConcurrent(clientId string, input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	// After validating inputs process
	result, err := s.CreateRepo(clientId, input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}

	output <- repositories.CreateRepositoriesResult{
		Response: result,
	}
}

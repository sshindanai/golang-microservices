package services

import (
	"strings"

	"github.com/sshindanai/golang-microservices/src/api/api/errors"
	"github.com/sshindanai/golang-microservices/src/api/config"
	"github.com/sshindanai/golang-microservices/src/api/domain/github"
	"github.com/sshindanai/golang-microservices/src/api/domain/providers/githib_provider"
	"github.com/sshindanai/golang-microservices/src/api/domain/repositories"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

// For mocking purpose
func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repository name")
	}
	req := github.CreateRepoRequest{
		Name:        input.Name,
		Private:     false,
		Description: input.Description,
	}

	// Small bug on GetAccessToken
	response, err := githib_provider.CreateRepo(config.GithubAccessToken, req)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}

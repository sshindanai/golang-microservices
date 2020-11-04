package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	GithubAccessToken    = "016cd1906270c87fc65da8c1724eb64ee9e45c35"
	LogLevel             = "info"
)

func IsProduction() bool {
	if os.Getenv("GO_ENVIRONMENT") == "production" {
		return true
	}
	return false
}

package oauth

import (
	"fmt"

	"github.com/sshindanai/golang-microservices/src/api/api/errors"
)

var (
	tokens = make(map[string]*AccessToken)
)

func (at *AccessToken) Save() errors.ApiError {
	// Generate access token
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserId)
	tokens[at.AccessToken] = at

	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.ApiError) {
	token := tokens[accessToken]
	if token == nil {
		return nil, errors.NewNotFoundError("no access token found with given parameter")
	}

	return token, nil
}

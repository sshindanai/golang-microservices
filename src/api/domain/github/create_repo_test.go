package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoCreateJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "golang introduction repository",
		Homepage:    "https://github.com",
		private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	// Action
	bytes, err := json.Marshal(request)

	// Validation
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreateRepoRequest

	// Unmarshal takes an input byte array and a *pointer* that we're trying to fill using this json
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
}

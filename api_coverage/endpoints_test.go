package api_coverage

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	leaders2 "github.com/will7200/go-wakatime/client/leaders"
	"github.com/will7200/go-wakatime/client/user"
	"testing"
	"time"
)

func TestClientAuthorization(t *testing.T) {
	_, err := client.User.Commits(nil, nil)
	assert.NotNil(t, err, "Client Authorization should have occurred")

	_, err = client.User.Duration(nil, nil)
	assert.NotNil(t, err, "Client Authorization should have occurred")

	_, err = client.User.Goals(nil, nil)
	assert.NotNil(t, err, "Client Authorization should have occurred")

	_, err = client.User.Projects(nil, nil)
	assert.NotNil(t, err, "Client Authorization should have occurred")

	_, err = client.User.Summaries(nil, nil)
	assert.NotNil(t, err, "Client Authorization should have occurred")

	_, _, err = client.User.Stats(nil, nil)
	assert.NotNil(t, err, "Client Authorization should have occurred")

	_, err = client.User.UserAgents(nil, nil)
	assert.NotNil(t, err, "Client Authorization should have occurred")
}

func TestUserCommitsEndpoint(t *testing.T) {
	params := user.NewCommitsParams()
	commits, err := client.User.Commits(params, apiKeyAuth)

	assert.NotNil(t, err, "Err should be null, no project was given")
	assert.Nil(t, commits, "Commits should be nil")
}

func TestUserDurationEndpoint(t *testing.T) {
	params := user.NewDurationParams()
	durations, err := client.User.Duration(params, apiKeyAuth)

	assert.NotNil(t, err, "Err should be null, no date parameter was passed")
	assert.Nil(t, durations, "Durations should not be nil")

	params = user.NewDurationParams()
	params.Date = strfmt.Date(time.Now().Add(time.Hour * 24))

	durations, err = client.User.Duration(params, apiKeyAuth)

	// If you suspect there is something wrong with the url request either set this to compare with the request being made
	// location := globalTransporter.urlRequests.Size()
	// fmt.Println(globalTransporter.urlRequests.Get(location - 1))

	assert.Nil(t, err, "Err should be null")
	assert.NotNil(t, durations, "Durations should not be nil")
}

func TestUserGoalsEndpoint(t *testing.T) {
	params := user.NewGoalsParams()
	goals, err := client.User.Goals(params, apiKeyAuth)

	assert.Nil(t, err, "Err should be null")
	assert.NotNil(t, goals, "Goals should not be nil")

}

func TestUserProjectsEndpoint(t *testing.T) {
	params := user.NewProjectsParams()
	projects, err := client.User.Projects(params, apiKeyAuth)

	assert.Nil(t, err, "Err should be null")
	assert.NotNil(t, projects, "Projects should not be nil")

}

func TestUserSummariesEndpoint(t *testing.T) {
	params := user.SummariesParams{
		Start:   strfmt.Date(time.Now().Add(-1 * time.Hour * 24)),
		End:     strfmt.Date(time.Now()),
		User:    "current",
		Context: context.Background(),
	}
	currentUser, err := client.User.Summaries(&params, apiKeyAuth)
	assert.Nil(t, err, "Err should be null")
	assert.NotNil(t, currentUser, "User Summaries should not be nil for the current user")
}

func TestUserStatsEndpoint(t *testing.T) {
	params := user.NewStatsParams()
	stats, ok, err := client.User.Stats(params, apiKeyAuth)

	assert.Nil(t, err, "Err should be null")
	assert.True(t, (stats == nil) != (ok == nil), "Should've received a positive response for stats")
}

func TestUserAgentsEndpoint(t *testing.T) {
	params := user.NewUserAgentsParams()
	durations, err := client.User.UserAgents(params, apiKeyAuth)

	assert.Nil(t, err, "Err should be null")
	assert.NotNil(t, durations, "Agents should not be nil")
}

func TestUserEndpoint(t *testing.T) {
	currentUser, err := client.User.User(nil, apiKeyAuth)

	assert.Nil(t, err, "Err should be null")
	assert.NotNil(t, currentUser, "User should not be nil for the current user")
}

func TestLeaderBoard(t *testing.T) {
	params := leaders2.NewLeaderParams()
	leaders, err := client.Leaders.Leader(params, nil)

	assert.Nil(t, err, "Err should be null")
	assert.NotNil(t, leaders, "Leaders should not be null, no api is needed")
}

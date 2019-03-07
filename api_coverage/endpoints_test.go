package api_coverage

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/will7200/go-wakatime/client/user"
	"testing"
	"time"
)

func TestSummariesEndpoint(t *testing.T) {
	params := user.SummariesParams{
		Start:   strfmt.Date(time.Now().Add(-1 * time.Hour * 24)),
		End:     strfmt.Date(time.Now()),
		User:    "current",
		Context: context.Background(),
	}
	currentUser, err := client.User.Summaries(&params, apiKeyAuth)
	assert.Nil(t, err, "Err should be null")
	assert.NotNil(t, currentUser, "User should not be nil for the current user")
}

func TestLeaderBoard(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {
	_, err := client.User.Summaries(nil, nil)
	assert.NotNil(t, err, "Client Authorization should have occured")
}

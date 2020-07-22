package wakatime

import (
	"context"
	"os"
	"testing"
)

func TestUserAgents(t *testing.T) {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	userID := os.Getenv("WAKATIME_USER_ID")

	client := NewClient(apiKey, nil)
	ctx := context.Background()
	query1 := &UserAgentsQuery{}
	_, err := client.UserAgents.Current(ctx, query1)
	if err != nil {
		t.Error(err)
	}

	_, err = client.UserAgents.User(ctx, userID, query1)
	if err != nil {
		t.Error(err)
	}

}

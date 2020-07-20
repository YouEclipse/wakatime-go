package wakatime

import (
	"context"
	"os"
	"testing"
)

func TestCommits(t *testing.T) {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	userID := os.Getenv("WAKATIME_USER_ID")
	projectID := os.Getenv("WAKATIME_PROJECT_ID")
	client := NewClient(apiKey, nil)
	ctx := context.Background()
	query := &CommitsQuery{}
	_, err := client.Commits.Current(ctx, projectID, query)
	if err != nil {
		t.Error(err)
	}

	_, err = client.Commits.User(ctx, userID, projectID, query)
	if err != nil {
		t.Error(err)
	}

}

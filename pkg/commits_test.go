package wakatime

import (
	"context"
	"os"
	"testing"
)

func TestCommits(t *testing.T) {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	client := NewClient(apiKey, nil)
	ctx := context.Background()
	query := &CommitsQuery{}
	_, err := client.Commits.Current(ctx, "59f75063-a117-44cc-a744-470398d682f2", query)
	if err != nil {
		t.Error(err)
	}

	_, err = client.Commits.User(ctx, "a5b4feda-214d-4ef2-bdc5-9a844c045006", "59f75063-a117-44cc-a744-470398d682f2", query)
	if err != nil {
		t.Error(err)
	}

}

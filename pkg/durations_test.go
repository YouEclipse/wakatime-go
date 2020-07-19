package wakatime

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestDurations(t *testing.T) {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	userID := os.Getenv("WAKATIME_USER_ID")

	client := NewClient(apiKey, nil)
	ctx := context.Background()
	query1 := &DurationsQuery{}
	_, err := client.Durations.Current(ctx, query1)
	if err != nil {
		t.Error(err)
	}
	query2 := &DurationsQuery{Date: String(time.Now().UTC().Format("2006-01-02"))}
	_, err = client.Durations.User(ctx, userID, query2)
	if err != nil {
		t.Error(err)
	}

}

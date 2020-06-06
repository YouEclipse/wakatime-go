package wakatime

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestDurations(t *testing.T) {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	client := NewClient(apiKey, nil)
	ctx := context.Background()
	query1 := &DurationsQuery{}
	_, err := client.Durations.Current(ctx, query1)
	if err != nil {
		t.Error(err)
	}
	query2 := &DurationsQuery{Date: String(time.Now().UTC().Format("2006-01-02"))}
	_, err = client.Durations.User(ctx, "a5b4feda-214d-4ef2-bdc5-9a844c045006", query2)
	if err != nil {
		t.Error(err)
	}

}

package wakatime

import (
	"context"
	"os"
	"testing"
)

func TestStats(t *testing.T) {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	client := NewClient(apiKey, nil)
	ctx := context.Background()
	query := &StatsQuery{}
	_, err := client.Stats.Current(ctx, RangeLast7Days, query)
	if err != nil {
		t.Error(err)
	}

	_, err = client.Stats.User(ctx, "a5b4feda-214d-4ef2-bdc5-9a844c045006", RangeLast7Days, query)
	if err != nil {
		t.Error(err)
	}

}

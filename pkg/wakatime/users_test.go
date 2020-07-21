package wakatime

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestUsers(t *testing.T) {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	userID := os.Getenv("WAKATIME_USER_ID")

	client := NewClient(apiKey, nil)
	ctx := context.Background()
	query1 := &UsersQuery{}
	resp, err := client.Users.Current(ctx, query1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", *resp.Data.ID)
	query2 := &UsersQuery{}
	_, err = client.Users.User(ctx, userID, query2)
	if err != nil {
		t.Error(err)
	}

}

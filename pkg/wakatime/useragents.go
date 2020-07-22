package wakatime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// UserAgentsService defines endpoint of the useragents API.
// A single user's profile information
type UserAgentsService service

// UserAgentsQuery defines the query of the userAgents API.
type UserAgentsQuery struct {
}

// UserAgentsResponse defines the response of the userAgents API.
// see https://wakatime.com/developers#userAgents for more information.
type UserAgentsResponse struct {
	Data *[]UserAgentsData `json:"data,omitempty"`
}

// UserAgentsData defines the data of the UserAgentsResponse.
type UserAgentsData struct {
	ID        *string `json:"id,omitempty"`
	Value     *string `json:"value,omitempty"`
	Editor    *string `json:"editor,omitempty"`
	Version   *string `json:"version,omitempty"`
	OS        *string `json:"os,omitempty"`
	LastSeen  *string `json:"last_seen,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
}

// Current do the request of the UserAgents API with current user.
func (u *UserAgentsService) Current(ctx context.Context, query *UserAgentsQuery) (*UserAgentsResponse, error) {
	return u.User(ctx, "current", query)
}

// User do the request of the UserAgents API with the given user.
func (u *UserAgentsService) User(ctx context.Context, userID string, query *UserAgentsQuery) (*UserAgentsResponse, error) {
	q := Query{}
	q.Parse(query)
	endpoint := fmt.Sprintf("users/%s/user_agents", userID)
	return u.do(ctx, endpoint, (url.Values)(q))
}

func (u *UserAgentsService) do(ctx context.Context, endpoint string, query url.Values) (*UserAgentsResponse, error) {
	respBytes, err := (*service)(u).get(ctx, endpoint, query)
	if err != nil {
		return nil, err
	}

	var response = &UserAgentsResponse{}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error  :%w", err)
	}

	return response, nil
}

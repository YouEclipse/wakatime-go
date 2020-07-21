package wakatime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// ProjectsService defines endpoint of the Projects API.
// List of WakaTime projects for a user.
type ProjectsService service

// ProjectsQuery defines the query of the Projects API.
type ProjectsQuery struct{}

// ProjectsResponse defines the response of the Projects API.
// see https://wakatime.com/developers#Projects for more information.
type ProjectsResponse struct {
	Data []Project `json:"data,omitempty"`
}

// Current do the request of the Projects API with current user.
func (p *ProjectsService) Current(ctx context.Context, query *ProjectsQuery) (*ProjectsResponse, error) {
	return p.User(ctx, "current", query)

}

// User do the request of the Projects API with the given user.
func (p *ProjectsService) User(ctx context.Context, userID string, query *ProjectsQuery) (*ProjectsResponse, error) {
	q := Query{}
	q.Parse(query)
	endpoint := fmt.Sprintf("users/%s/projects", userID)
	return p.do(ctx, endpoint, (url.Values)(q))
}

func (p *ProjectsService) do(ctx context.Context, endpoint string, query url.Values) (*ProjectsResponse, error) {
	respBytes, err := (*service)(p).get(ctx, endpoint, query)
	if err != nil {
		return nil, err
	}

	var response = &ProjectsResponse{}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error  :%w", err)
	}

	return response, nil
}

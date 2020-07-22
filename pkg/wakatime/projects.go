package wakatime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
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

// Project is the project of the commit.
type Project struct {
	CreatedAt                    *time.Time  `json:"created_at,omitempty"`
	HasPublicURL                 *bool       `json:"has_public_url,omitempty"`
	HTMLEscapedName              *string     `json:"html_escaped_name,omitempty"`
	HumanReadableLastHeartbeatAt *string     `json:"human_readable_last_heartbeat_at,omitempty"`
	ID                           *string     `json:"id,omitempty"`
	LastHeartbeatAt              *time.Time  `json:"last_heartbeat_at,omitempty"`
	Name                         *string     `json:"name,omitempty"`
	Repository                   *Repository `json:"repository,omitempty"`
	URL                          *string     `json:"url,omitempty"`
}

// Repository is the repository of the project.
type Repository struct {
	Badge               *string    `json:"badge,omitempty"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	DefaultBranch       *string    `json:"default_branch,omitempty"`
	Description         *string    `json:"description,omitempty"`
	ForkCount           *int64     `json:"fork_count,omitempty"`
	FullName            *string    `json:"full_name,omitempty"`
	Homepage            *string    `json:"homepage,omitempty"`
	HTMLURL             *string    `json:"html_url,omitempty"`
	ID                  *string    `json:"id,omitempty"`
	ImageIconURL        *string    `json:"image_icon_url,omitempty"`
	IsFork              *bool      `json:"is_fork,omitempty"`
	IsPrivate           *bool      `json:"is_private,omitempty"`
	LastSyncedAt        *string    `json:"last_synced_at,omitempty"`
	ModifiedAt          *time.Time `json:"modified_at,omitempty"`
	Name                *string    `json:"name,omitempty"`
	Provider            *string    `json:"provider,omitempty"`
	StarCount           *int64     `json:"star_count,omitempty"`
	URL                 *string    `json:"url,omitempty"`
	WakatimeProjectName *string    `json:"wakatime_project_name,omitempty"`
	WatchCount          *int64     `json:"watch_count,omitempty"`
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

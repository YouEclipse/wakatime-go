package wakatime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// CommitService defines endpoint of the commits API.
// List of commits for a WakaTime project showing the time spent coding in each commit.
type CommitService service

// CommitsQuery defines the query of the commits API.
type CommitsQuery struct {
	Author *string `query:"author"`
	Page   *int64  `query:"page"`
}

// CommitsResponse defines the response of the commits API.
// see https://wakatime.com/developers#commits for more information.
type CommitsResponse struct {
	Author      *string   `json:"author,omitempty"`
	Commits     []*Commit `json:"commits,omitempty"`
	NextPage    *int64    `json:"next_page,omitempty"`
	NextPageURL *string   `json:"next_page_url,omitempty"`
	Page        int       `json:"page,omitempty"`
	PrevPage    *int64    `json:"prev_page,omitempty"`
	PrevPageURL *string   `json:"prev_page_url,omitempty"`
	Project     *Project  `json:"project,omitempty"`
	Status      *string   `json:"status,omitempty"`
	TotalPages  int       `json:"total_pages,omitempty"`
}

// Commit defines the commit data of the commits API response.
type Commit struct {
	AuthorAvatarURL               *string    `json:"author_avatar_url,omitempty"`
	AuthorDate                    *time.Time `json:"author_date,omitempty"`
	AuthorEmail                   *string    `json:"author_email,omitempty"`
	AuthorHTMLURL                 *string    `json:"author_html_url,omitempty"`
	AuthorID                      *string    `json:"author_id,omitempty"`
	AuthorName                    *string    `json:"author_name,omitempty"`
	AuthorURL                     *string    `json:"author_url,omitempty"`
	AuthorUsername                *string    `json:"author_username,omitempty"`
	Branch                        *string    `json:"branch,omitempty"`
	CommitterAvatarURL            *string    `json:"committer_avatar_url,omitempty"`
	CommitterDate                 *time.Time `json:"committer_date,omitempty"`
	CommitterEmail                *string    `json:"committer_email,omitempty"`
	CommitterHTMLURL              *string    `json:"committer_html_url,omitempty"`
	CommitterName                 *string    `json:"committer_name,omitempty"`
	CommitterURL                  *string    `json:"committer_url,omitempty"`
	CommitterUsername             *string    `json:"committer_username,omitempty"`
	CreatedAt                     *time.Time `json:"created_at,omitempty"`
	Hash                          *string    `json:"hash,omitempty"`
	HTMLURL                       *string    `json:"html_url,omitempty"`
	HumanReadableDate             *string    `json:"human_readable_date,omitempty"`
	HumanReadableNaturalDate      *string    `json:"human_readable_natural_date,omitempty"`
	HumanReadableTotal            *string    `json:"human_readable_total,omitempty"`
	HumanReadableTotalWithSeconds *string    `json:"human_readable_total_with_seconds,omitempty"`
	ID                            *string    `json:"id,omitempty"`
	Message                       *string    `json:"message,omitempty"`
	Ref                           *string    `json:"ref,omitempty"`
	TotalSeconds                  *int64     `json:"total_seconds,omitempty"`
	TruncatedHash                 *string    `json:"truncated_hash,omitempty"`
	URL                           *string    `json:"url,omitempty"`
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

// Current do the request of the commits API with current user.
func (d *CommitService) Current(ctx context.Context, projectID string, query *CommitsQuery) (*CommitsResponse, error) {
	return d.User(ctx, "current", projectID, query)

}

// User do the request of the commits API with the given user.
func (d *CommitService) User(ctx context.Context, userID, projectID string, query *CommitsQuery) (*CommitsResponse, error) {
	endpoint := fmt.Sprintf("users/%s/projects/%s/commits", userID, projectID)

	q := Query{}
	q.Parse(query)
	return d.do(ctx, endpoint, (url.Values)(q))
}

func (d *CommitService) do(ctx context.Context, endpoint string, query url.Values) (*CommitsResponse, error) {
	respBytes, err := (*service)(d).get(ctx, endpoint, query)
	if err != nil {
		return nil, err
	}

	var response = &CommitsResponse{}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error  :%w", err)
	}

	return response, nil
}

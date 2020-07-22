package wakatime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// UsersService defines endpoint of the users API.
// A single user's profile information
type UsersService service

// UsersQuery defines the query of the users API.
type UsersQuery struct {
}

// UsersResponse defines the response of the users API.
// see https://wakatime.com/developers#users for more information.
type UsersResponse struct {
	Data *UsersData `json:"data,omitempty"`
}

// UsersData defines the data of the UsersResponse.
type UsersData struct {
	ID                   *string `json:"id,omitempty"`
	HasPremiumFeatures   *bool   `json:"has_premium_features,omitempty"`
	DisplayName          *string `json:"display_name,omitempty"`
	FullName             *string `json:"full_name,omitempty"`
	Email                *string `json:"email,omitempty"`
	Photo                *string `json:"photo,omitempty"`
	IsEmailPublic        *bool   `json:"is_email_public,omitempty"`
	IsEmailConfirmed     *bool   `json:"is_email_confirmed,omitempty"`
	PublicEmail          *string `json:"public_email,omitempty"`
	PhotoPublic          *bool   `json:"photo_public,omitempty"`
	Timezone             *string `json:"timezone,omitempty"`
	LastHeartbeatAt      *string `json:"last_heartbeat_at,omitempty"`
	LastPlugin           *string `json:"last_plugin,omitempty"`
	LastPluginName       *string `json:"last_plugin_name,omitempty"`
	LastProject          *string `json:"last_project,omitempty"`
	Plan                 *string `json:"plan,omitempty"`
	Username             *string `json:"username,omitempty"`
	Website              *string `json:"website,omitempty"`
	HumanReadableWebsite *string `json:"human_readable_website,omitempty"`
	Location             *string `json:"location,omitempty"`
	LoggedTimePublic     *bool   `json:"logged_time_public,omitempty"`
	LanguagesUsedPublic  *bool   `json:"languages_used_public,omitempty"`
	IsHireable           *bool   `json:"is_hireable,omitempty"`
	CreatedAt            *string `json:"created_at,omitempty"`
	ModifiedAt           *string `json:"modified_at,omitempty"`
}

// Current do the request of the Users API with current user.
func (u *UsersService) Current(ctx context.Context, query *UsersQuery) (*UsersResponse, error) {
	return u.User(ctx, "current", query)
}

// User do the request of the Users API with the given user.
func (u *UsersService) User(ctx context.Context, userID string, query *UsersQuery) (*UsersResponse, error) {
	q := Query{}
	q.Parse(query)
	endpoint := fmt.Sprintf("users/%s/", userID)
	return u.do(ctx, endpoint, (url.Values)(q))
}

func (u *UsersService) do(ctx context.Context, endpoint string, query url.Values) (*UsersResponse, error) {
	respBytes, err := (*service)(u).get(ctx, endpoint, query)
	if err != nil {
		return nil, err
	}

	var response = &UsersResponse{}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error  :%w", err)
	}

	return response, nil
}

package wakatime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// GoalsService defines endpoint of the Goals API.
// A user's coding activity for the given day as an array of Goals.
type GoalsService service

// GoalsQuery defines the query of the Goals API.
type GoalsQuery struct{}

// GoalsResponse defines the response of the Goals API.
// see https://wakatime.com/developers#Goals for more information.
type GoalsResponse struct {
	Data       []*GoalsData `json:"data,omitempty"`
	Total      *int         `json:"total,omitempty"`
	TotalPages *int         `json:"total_pages,omitempty"`
}

// GoalsData defines the data of the GoalsResponse.
type GoalsData struct {
	AverageStatus    *string        `json:"average_status,omitempty"`
	ChartData        *ChartData     `json:"chart_data,omitempty"`
	CumulativeStatus *string        `json:"cumulative_status,omitempty"`
	Delta            *string        `json:"delta,omitempty"`
	ID               *string        `json:"id,omitempty"`
	IgnoreDays       []*string      `json:"ignore_days,omitempty"`
	ImproveByPercent *float64       `json:"improve_by_percent,omitempty"`
	IsEnabled        *bool          `json:"is_enabled,omitempty"`
	Languages        []*string      `json:"languages,omitempty"`
	Projects         []*string      `json:"projects,omitempty"`
	RangeText        *string        `json:"range_text,omitempty"`
	Seconds          *int           `json:"seconds,omitempty"`
	Status           *string        `json:"status,omitempty"`
	Subscribers      []*Subscribers `json:"subscribers,omitempty"`
	Title            *string        `json:"title,omitempty"`
	Type             *string        `json:"type,omitempty"`
}

// ChartData defines the data of the Chart of the GoalsData
type ChartData struct {
	ActualSeconds     *float64   `json:"actual_seconds,omitempty"`
	ActualSecondsText *string    `json:"actual_seconds_text,omitempty"`
	GoalSeconds       *int       `json:"goal_seconds,omitempty"`
	GoalSecondsText   *string    `json:"goal_seconds_text,omitempty"`
	Range             *RangeData `json:"range,omitempty"`
}

// RangeData defines the Range of the ChartData
type RangeData struct {
	Date     *string `json:"date,omitempty"`
	End      *string `json:"end,omitempty"`
	Start    *string `json:"start,omitempty"`
	Text     *string `json:"text,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// Subscribers defines the Subscribers of the GoalsData
type Subscribers struct {
	Email          *string `json:"email,omitempty"`
	EmailFrequency *string `json:"email_frequency,omitempty"`
	FullName       *string `json:"full_name,omitempty"`
	UserID         *string `json:"user_id,omitempty"`
	Username       *string `json:"username,omitempty"`
}

// Current do the request of the goals API with current user.
func (g *GoalsService) Current(ctx context.Context, query *GoalsQuery) (*GoalsResponse, error) {
	return g.User(ctx, "current", query)

}

// User do the request of the Goals API with the given user.
func (g *GoalsService) User(ctx context.Context, userID string, query *GoalsQuery) (*GoalsResponse, error) {
	q := Query{}
	q.Parse(query)
	endpoint := fmt.Sprintf("users/%s/goals", userID)
	return g.do(ctx, endpoint, (url.Values)(q))
}

func (g *GoalsService) do(ctx context.Context, endpoint string, query url.Values) (*GoalsResponse, error) {
	respBytes, err := (*service)(g).get(ctx, endpoint, query)
	if err != nil {
		return nil, err
	}

	var response = &GoalsResponse{}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error  :%w", err)
	}

	return response, nil
}

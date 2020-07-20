package wakatime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// DurationService defines endpoint of the durations API.
// A user's coding activity for the given day as an array of durations.
type DurationService service

// DurationsQuery defines the query of the durations API.
type DurationsQuery struct {
	Date       *string `query:"date"`
	Project    *string `query:"project"`
	Branches   *string `query:"branches"`
	Timeouts   *string `query:"timeouts"`
	WritesOnly *bool   `query:"writes_only"`
}

// DurationsResponse defines the response of the durations API.
// see https://wakatime.com/developers#durations for more information.
type DurationsResponse struct {
	Branches []*string        `json:"branches,omitempty"`
	Data     []*DurationsData `json:"data,omitempty"`
	End      *time.Time       `json:"end,omitempty"`
	Start    *time.Time       `json:"start,omitempty"`
	Timezone *string          `json:"timezone,omitempty"`
}

// DurationsData defines the data of the DurationsResponse.
type DurationsData struct {
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	Cursorpos     *int64     `json:"cursorpos,omitempty"`
	Duration      *float64   `json:"duration,omitempty"`
	ID            *string    `json:"id,omitempty"`
	Lineno        *int64     `json:"lineno,omitempty"`
	MachineNameID *string    `json:"machine_name_id,omitempty"`
	Project       *string    `json:"project,omitempty"`
	Time          *float64   `json:"time,omitempty"`
	UserID        *string    `json:"user_id,omitempty"`
}

// Current do the request of the durations API with current user.
func (d *DurationService) Current(ctx context.Context, query *DurationsQuery) (*DurationsResponse, error) {
	return d.User(ctx, "current", query)

}

// User do the request of the durations API with the given user.
func (d *DurationService) User(ctx context.Context, userID string, query *DurationsQuery) (*DurationsResponse, error) {
	if query.Date == nil || *query.Date == "" {
		query.Date = String(time.Now().UTC().Format("2006-01-02"))
	}

	q := Query{}
	q.Parse(query)
	endpoint := fmt.Sprintf("users/%s/durations", userID)
	return d.do(ctx, endpoint, (url.Values)(q))
}

func (d *DurationService) do(ctx context.Context, endpoint string, query url.Values) (*DurationsResponse, error) {
	respBytes, err := (*service)(d).get(ctx, endpoint, query)
	if err != nil {
		return nil, err
	}

	var response = &DurationsResponse{}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error  :%w", err)
	}

	return response, nil
}

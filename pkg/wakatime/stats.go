package wakatime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// StatService defines endpoint of the durations API.
// A user's coding activity for the given time range.
type StatService service

// TimeRange defines the string type of time range.
type TimeRange string

const (
	// RangeLast7Days is the range of last 7 days
	RangeLast7Days TimeRange = "last_7_days"
	// RangeLast30Days is the range of last 30 days
	RangeLast30Days TimeRange = "last_30_days"
	// RangeLast6Months is the range of last 6 months
	RangeLast6Months TimeRange = "last_6_months"
	// RangeLastYear is the range of last year
	RangeLastYear TimeRange = "last_year"
)

// StatsQuery defines query of the stats API.
type StatsQuery struct {
	Timeout    *int64  `query:"timeout"`
	WritesOnly *bool   `query:"writes_only"`
	Project    *string `query:"project"`
}

// StatsResponse defines the response of the stats API.
type StatsResponse struct {
	Data StatsData `json:"data,omitempty"`
}

// StatsData defines the data of stats API response.
type StatsData struct {
	BestDay                                         *BestDay      `json:"best_day,omitempty"`
	Categories                                      []*StatItem   `json:"categories,omitempty"`
	CreatedAt                                       *time.Time    `json:"created_at,omitempty"`
	DailyAverage                                    *float64      `json:"daily_average,omitempty"`
	DailyAverageIncludingOtherLanguage              *float64      `json:"daily_average_including_other_language,omitempty"`
	DaysIncludingHolidays                           *int64        `json:"days_including_holidays,omitempty"`
	DaysMinusHolidays                               *int64        `json:"days_minus_holidays,omitempty"`
	Dependencies                                    []interface{} `json:"dependencies,omitempty"`
	Editors                                         []*StatItem   `json:"editors,omitempty"`
	End                                             *time.Time    `json:"end,omitempty"`
	Holidays                                        *int64        `json:"holidays,omitempty"`
	HumanReadableDailyAverage                       *string       `json:"human_readable_daily_average,omitempty"`
	HumanReadableDailyAverageIncludingOtherLanguage *string       `json:"human_readable_daily_average_including_other_language,omitempty"`
	HumanReadableTotal                              *string       `json:"human_readable_total,omitempty"`
	HumanReadableTotalIncludingOtherLanguage        *string       `json:"human_readable_total_including_other_language,omitempty"`
	ID                                              *string       `json:"id,omitempty"`
	IsAlreadyUpdating                               *bool         `json:"is_already_updating,omitempty"`
	IsCodingActivityVisible                         *bool         `json:"is_coding_activity_visible,omitempty"`
	IsIncludingToday                                *bool         `json:"is_including_today,omitempty"`
	IsOtherUsageVisible                             *bool         `json:"is_other_usage_visible,omitempty"`
	IsStuck                                         *bool         `json:"is_stuck,omitempty"`
	IsUpToDate                                      *bool         `json:"is_up_to_date,omitempty"`
	Languages                                       []StatItem    `json:"languages,omitempty"`
	Machines                                        []Machines    `json:"machines,omitempty"`
	ModifiedAt                                      *time.Time    `json:"modified_at,omitempty"`
	OperatingSystems                                []*StatItem   `json:"operating_systems,omitempty"`
	Project                                         *interface{}  `json:"project,omitempty"`
	Projects                                        []*StatItem   `json:"projects,omitempty"`
	Range                                           *string       `json:"range,omitempty"`
	Start                                           *time.Time    `json:"start,omitempty"`
	Status                                          *string       `json:"status,omitempty"`
	Timeout                                         *int64        `json:"timeout,omitempty"`
	Timezone                                        *string       `json:"timezone,omitempty"`
	TotalSeconds                                    *float64      `json:"total_seconds,omitempty"`
	TotalSecondsIncludingOtherLanguage              *float64      `json:"total_seconds_including_other_language,omitempty"`
	UserID                                          *string       `json:"user_id,omitempty"`
	Username                                        *string       `json:"username,omitempty"`
	WritesOnly                                      *bool         `json:"writes_only,omitempty"`
}

// BestDay is the day with most coding time.
type BestDay struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	Date         *string    `json:"date,omitempty"`
	ID           *string    `json:"id,omitempty"`
	ModifiedAt   *time.Time `json:"modified_at,omitempty"`
	Text         *string    `json:"text,omitempty"`
	TotalSeconds *float64   `json:"total_seconds,omitempty"`
}

// StatItem is the item for a stat.
type StatItem struct {
	Digital      *string  `json:"digital,omitempty"`
	Hours        *int64   `json:"hours,omitempty"`
	Minutes      *int64   `json:"minutes,omitempty"`
	Name         *string  `json:"name,omitempty"`
	Percent      *float64 `json:"percent,omitempty"`
	Text         *string  `json:"text,omitempty"`
	TotalSeconds *float64 `json:"total_seconds,omitempty"`
}

// Machine is the details of machine.
type Machine struct {
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ID         *string    `json:"id,omitempty"`
	IP         *string    `json:"ip,omitempty"`
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`
	Name       *string    `json:"name,omitempty"`
	Value      *string    `json:"value,omitempty"`
}

// Machines is the stats for machines.
type Machines struct {
	Digital      *string  `json:"digital,omitempty"`
	Hours        *int64   `json:"hours,omitempty"`
	Machine      *Machine `json:"machine,omitempty"`
	Minutes      *int64   `json:"minutes,omitempty"`
	Name         *string  `json:"name,omitempty"`
	Percent      *float64 `json:"percent,omitempty"`
	Text         *string  `json:"text,omitempty"`
	TotalSeconds *float64 `json:"total_seconds,omitempty"`
}

// Current do the request of the durations API with current user.
func (d *StatService) Current(ctx context.Context, timeRange TimeRange, query *StatsQuery) (*StatsResponse, error) {
	return d.User(ctx, "current", timeRange, query)

}

// User do the request of the durations API with the given user.
func (d *StatService) User(ctx context.Context, userID string, timeRange TimeRange, query *StatsQuery) (*StatsResponse, error) {

	q := Query{}
	q.Parse(query)
	endpoint := fmt.Sprintf("users/%s/stats/%s", userID, timeRange)
	return d.do(ctx, endpoint, (url.Values)(q))
}

func (d *StatService) do(ctx context.Context, endpoint string, query url.Values) (*StatsResponse, error) {
	respBytes, err := (*service)(d).get(ctx, endpoint, query)
	if err != nil {
		return nil, err
	}

	var response = &StatsResponse{}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error  :%w", err)
	}

	return response, nil
}

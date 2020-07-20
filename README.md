# wakatime-go

[WIP] 🕘 Go library for accessing the [Wakatime](https://wakatime.com/developers#introduction) API

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/YouEclipse/wakatime-go/pkg) ![Go](https://github.com/YouEclipse/wakatime-go/workflows/Go/badge.svg)

## Install

#### Requirments

> Go version 1.13+

#### Install

```
go get github.com/YouEclipse/wakatime-go
```

## Quick start

```golang

import (
	"github.com/YouEclipse/wakatime-go/pkg/wakatime"
)

func main() {

	apiKey := os.Getenv("WAKATIME_API_KEY")
	client := wakatime.NewClient(apiKey, &http.Client{})

	ctx := context.Background()
	query := &wakatime.StatsQuery{}

	stats, err := client.Stats.Current(ctx, wakatime.RangeLast7Days, query)

    ...
}



```

## Features v0.1.0

#### TODOs

- [x] [Commits](https://wakatime.com/developers#commits)
- [x] [Durations](https://wakatime.com/developers#durations)
- [ ] [Heartbeats](https://wakatime.com/developers#heartbeats)
- [ ] [Leaders](https://wakatime.com/developers#leaders)
- [ ] [Meta](https://wakatime.com/developers#meta)
- [ ] [Org Dashboard Member Durations](https://wakatime.com/developers#org_dashboard_member_durations)
- [ ] [Org Dashboard Member Summaries](https://wakatime.com/developers#org_dashboard_member_summaries)
- [ ] [Org Dashboard Members](https://wakatime.com/developers#org_dashboard_members)
- [ ] [Org Dashboards](https://wakatime.com/developers#org_dashboards)
- [ ] [Orgs](https://wakatime.com/developers#orgs)
- [ ] [Private Leaderboards](https://wakatime.com/developers#private_leaderboards)
- [ ] [Private Leaderboards Leaders](https://wakatime.com/developers#private_leaderboards_leaders)
- [ ] [Projects](https://wakatime.com/developers#projects)
- [x] [Stats](https://wakatime.com/developers#stats)
- [ ] [Summaries](https://wakatime.com/developers#summaries)
- [ ] [User Agents](https://wakatime.com/developers#user_agents)
- [ ] [Users](https://wakatime.com/developers#users)

...

## License

[Apache 2.0](./LICENSE)

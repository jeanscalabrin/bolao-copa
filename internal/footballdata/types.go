package footballdata

import "time"

type MatchesResponse struct {
	Matches []MatchResponse `json:"matches"`
}

type MatchResponse struct {
	ID        int           `json:"id"`
	UtcDate   time.Time     `json:"utcDate"`
	Status    string        `json:"status"`
	Stage     string        `json:"stage"`
	GroupName string        `json:"group"`
	HomeTeam  TeamResponse  `json:"homeTeam"`
	AwayTeam  TeamResponse  `json:"awayTeam"`
	Score     ScoreResponse `json:"score"`
}

type TeamResponse struct {
	Name string `json:"name"`
}

type ScoreResponse struct {
	Fulltime FullTimeScore `json:"fullTime"`
}

type FullTimeScore struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}

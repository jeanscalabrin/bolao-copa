package matches

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID         uuid.UUID
	ExternalID int
	HomeTeam   string
	AwayTeam   string
	Stage      string
	GroupName  string
	Status     string
	HomeScore  *int
	AwayScore  *int
	MatchDate  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

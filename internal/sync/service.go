package sync

import (
	"context"

	"github.com/google/uuid"
	"github.com/jeanscalabrin/bolao-copa/internal/footballdata"
	"github.com/jeanscalabrin/bolao-copa/internal/matches"
)

type Service struct {
	client *footballdata.Client
	repo   *matches.Repository
}

func NewService(client *footballdata.Client, repo *matches.Repository) *Service {
	return &Service{
		client: client,
		repo:   repo,
	}
}

func (s *Service) SyncMatches(ctx context.Context) error {
	apiMatches, err := s.client.GetWorldCupMatches()

	if err != nil {
		return err
	}

	for _, m := range apiMatches {
		match := matches.Match{
			ID:         uuid.New(),
			ExternalID: m.ID,
			HomeTeam:   m.HomeTeam.Name,
			AwayTeam:   m.AwayTeam.Name,
			Stage:      m.Stage,
			GroupName:  m.GroupName,
			Status:     m.Status,
			HomeScore:  m.Score.Fulltime.Home,
			AwayScore:  m.Score.Fulltime.Away,
			MatchDate:  m.UtcDate,
		}

		if err := s.repo.Upsert(ctx, match); err != nil {
			return err
		}
	}
	return nil
}

package matches

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List(ctx context.Context) ([]Match, error) {
	rows, err := r.db.Query(ctx,
		`SELECT
			id,
			home_team,
			away_team,
			match_date
		FROM matches
		ORDER BY match_date`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var matches []Match

	for rows.Next() {
		var m Match

		err := rows.Scan(
			&m.ID,
			&m.HomeTeam,
			&m.AwayTeam,
			&m.MatchDate,
		)

		if err != nil {
			return nil, err
		}

		matches = append(matches, m)

	}
	return matches, nil
}

func (r *Repository) Count(ctx context.Context) (int, error) {
	var count int

	err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM matches`).Scan(&count)

	return count, err
}

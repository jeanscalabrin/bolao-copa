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

func (r *Repository) Upsert(
	ctx context.Context,
	match Match,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO matches (
			id,
			external_id,
			home_team,
			away_team,
			status,
			home_score,
			away_score,
			match_date,
			group_name,
			stage
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10
		)
		ON CONFLICT (external_id)
		DO UPDATE SET
			group_name = EXCLUDED.group_name,
			stage = EXCLUDED.stage,
			status = EXCLUDED.status,
			home_score = EXCLUDED.home_score,
			away_score = EXCLUDED.away_score,
			match_date = EXCLUDED.match_date
		`,
		match.ID,
		match.ExternalID,
		match.HomeTeam,
		match.AwayTeam,
		match.Status,
		match.HomeScore,
		match.AwayScore,
		match.MatchDate,
		match.GroupName,
		match.Stage,
	)

	return err
}

func (r *Repository) List(ctx context.Context) ([]Match, error) {
	rows, err := r.db.Query(ctx,
		`SELECT
			id,
			home_team,
			away_team,
			home_score,
			away_score,
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
			&m.HomeScore,
			&m.AwayScore,
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

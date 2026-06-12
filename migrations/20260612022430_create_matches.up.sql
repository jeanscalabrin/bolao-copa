CREATE TABLE matches (
    id UUID PRIMARY KEY,
    external_id BIGINT NOT NULL UNIQUE,
    home_team TEXT NOT NULL,
    away_team TEXT NOT NULL,
    status TEXT NOT NULL,
    home_score INTEGER,
    away_score INTEGER,
    match_date TIMESTAMP NOT NULL,
    group TEXT NOT NULL,
    stage TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

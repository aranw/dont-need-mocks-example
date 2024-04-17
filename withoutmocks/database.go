package withoutmocks

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Score struct {
	UserID    int
	HighScore int
	Rank      int
}

func GetLeaderboard(ctx context.Context, conn *pgx.Conn) ([]Score, error) {
	var scores []Score
	rows, err := conn.Query(ctx, `SELECT
							t.user_id,
							t.high_score,
							RANK() OVER (ORDER BY t.high_score DESC) AS rank
						FROM (
							SELECT
							user_id,
							MAX(score) AS high_score
							FROM scores
							GROUP BY user_id
						) AS t
						ORDER BY rank`)
	if err != nil {
		return nil, fmt.Errorf("querying leaderboard: %w", err)
	}

	for rows.Next() {
		var score Score
		err := rows.Scan(&score.UserID, &score.HighScore, &score.Rank)
		if err != nil {
			return nil, fmt.Errorf("scanning row into score: %w", err)
		}
		scores = append(scores, score)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("scanning rows: %w", err)
	}

	return scores, nil
}

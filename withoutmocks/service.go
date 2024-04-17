package withoutmocks

import (
	"context"
	"fmt"
	"log/slog"
)

type Service struct {
	DB *Database
}

func (s *Service) GetLeaderboard() []string {
	scores, err := s.DB.Get(context.Background())
	if err != nil {
		slog.Error("something went wrong", "err", err)
		return nil
	}

	sc := make([]string, 0, len(scores))
	for _, score := range scores {
		formattedScore := fmt.Sprintf("User ID: %d, High Score: %d, Rank: %d", score.UserID, score.HighScore, score.Rank)
		sc = append(sc, formattedScore)
	}

	return sc
}

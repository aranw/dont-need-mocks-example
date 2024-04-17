package withmocks

import (
	"context"
	"fmt"
	"log/slog"
)

//go:generate mockgen -destination=./mocks/retrieve_leaderboard_mock.go -package=mocks . LeaderboardRetriever
type LeaderboardRetriever interface {
	Get(ctx context.Context) ([]Score, error)
}

type Service struct {
	LBR LeaderboardRetriever
}

func (s *Service) GetLeaderboard() []string {
	scores, err := s.LBR.Get(context.Background())
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

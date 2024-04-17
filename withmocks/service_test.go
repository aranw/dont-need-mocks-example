package withmocks_test

import (
	"context"
	"testing"

	"github.com/aranw/dont-need-mocks-example/withmocks"
	"github.com/aranw/dont-need-mocks-example/withmocks/mocks"
	"github.com/go-quicktest/qt"
	"go.uber.org/mock/gomock"
)

func TestGetLeaderboard(t *testing.T) {
	gomock := gomock.NewController(t)
	defer gomock.Finish()

	m := mocks.NewMockLeaderboardRetriever(gomock)

	m.EXPECT().Get(context.Background()).Times(1).Return([]withmocks.Score{
		{UserID: 1, HighScore: 5, Rank: 1},
		{UserID: 2, HighScore: 3, Rank: 2},
	}, nil)

	s := withmocks.Service{
		LBR: m,
	}

	scores := s.GetLeaderboard()

	qt.Assert(t, qt.HasLen(scores, 2))
	qt.Assert(t, qt.Equals(scores[0], "User ID: 1, High Score: 5, Rank: 1"))
	qt.Assert(t, qt.Equals(scores[1], "User ID: 2, High Score: 3, Rank: 2"))
}

package withmocks_test

import (
	"context"
	"testing"

	"github.com/aranw/dont-need-mocks-example/withmocks"
	"github.com/aranw/dont-need-mocks-example/withmocks/mocks"

	"github.com/go-quicktest/qt"
	"github.com/jackc/pgx/v5"
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

	scores, err := withmocks.GetLeaderboard(m)
	if err != nil {
		t.Fatal(err)
	}

	// User 1 is in Rank 1
	qt.Assert(t, qt.Equals(scores[0].HighScore, 5))
	qt.Assert(t, qt.Equals(scores[0].UserID, 1))
	qt.Assert(t, qt.Equals(scores[0].Rank, 1))

	// User 2 is in Rank 2
	qt.Assert(t, qt.Equals(scores[1].HighScore, 3))
	qt.Assert(t, qt.Equals(scores[1].UserID, 2))
	qt.Assert(t, qt.Equals(scores[1].Rank, 2))
}

func TestGetLeaderboard_PGX(t *testing.T) {
	// Use default connection settings for example
	conn, err := pgx.Connect(context.Background(), "")
	if err != nil {
		t.Fatalf("connecting to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	if _, err := conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS "scores" (
		"id" SERIAL PRIMARY KEY,
		"user_id" INTEGER,
		"score" INTEGER);`); err != nil {
		t.Fatal(err)
	}

	if _, err := conn.Exec(context.Background(), `TRUNCATE TABLE scores;`); err != nil {
		t.Fatal(err)
	}

	if _, err := conn.Exec(context.Background(), `
		INSERT INTO scores (user_id, score)
			VALUES 
				(1, 5),
				(1, 4),
				(2, 3),
				(2, 2);
	`); err != nil {
		t.Fatal(err)
	}

	pgxRetriever := &withmocks.PGXLeaderboardRetrieveer{
		Conn: conn,
	}

	scores, err := withmocks.GetLeaderboard(pgxRetriever)
	if err != nil {
		t.Fatal(err)
	}

	// User 1 is in Rank 1
	qt.Assert(t, qt.Equals(scores[0].HighScore, 5))
	qt.Assert(t, qt.Equals(scores[0].UserID, 1))
	qt.Assert(t, qt.Equals(scores[0].Rank, 1))

	// User 2 is in Rank 2
	qt.Assert(t, qt.Equals(scores[1].HighScore, 3))
	qt.Assert(t, qt.Equals(scores[1].UserID, 2))
	qt.Assert(t, qt.Equals(scores[1].Rank, 2))
}

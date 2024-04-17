package withoutmocks

import (
	"context"
	"testing"

	"github.com/go-quicktest/qt"
	"github.com/jackc/pgx/v5"
)

func TestService_GetLeaderboard(t *testing.T) {
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

	db := &Database{Conn: conn}

	service := &Service{
		DB: db,
	}

	scores := service.GetLeaderboard()
	qt.Assert(t, qt.HasLen(scores, 2))
	qt.Assert(t, qt.Equals(scores[0], "User ID: 1, High Score: 5, Rank: 1"))
	qt.Assert(t, qt.Equals(scores[1], "User ID: 2, High Score: 3, Rank: 2"))
}

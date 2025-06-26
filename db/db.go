package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("GOOSE_DBSTRING"))
	return conn, err
}

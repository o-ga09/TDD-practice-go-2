package db

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
)

func Connect(ctx context.Context) *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	slog.Log(ctx, slog.LevelInfo, "db connected")
	return db
}

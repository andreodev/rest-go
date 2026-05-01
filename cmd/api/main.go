package main

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"rest-go/internal/handlers"
	"rest-go/internal/repositories"
	"rest-go/internal/usecases"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	databaseURL := getDatabaseURL()

	migrationDB := openDatabase(databaseURL)
	runMigrations(migrationDB)
	migrationDB.Close()

	db := openDatabase(databaseURL)
	defer db.Close()

	repos := repositories.New(db)

	useCases := usecases.New(repos)

	h := handlers.New(useCases)

	if err := h.Listen(8080); err != nil {
		slog.Error("failed to listen", "err", err)
		os.Exit(1)
	}
}

func getDatabaseURL() string {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		slog.Error("DATABASE_URL is required")
		os.Exit(1)
	}

	return databaseURL
}

func openDatabase(databaseURL string) *sql.DB {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		slog.Error("failed to open database", "err", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		slog.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}

	slog.Info("database connected")

	return db
}

func runMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.Error("failed to create migration driver", "err", err)
		os.Exit(1)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		slog.Error("failed to load migrations", "err", err)
		os.Exit(1)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		slog.Error("failed to run migrations", "err", err)
		os.Exit(1)
	}

	slog.Info("migrations applied")
}

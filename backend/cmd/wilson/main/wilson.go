package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/JP-Go/wilson/backend/internal/application/api"
	"github.com/JP-Go/wilson/backend/internal/application/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()
	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("Bye")
}

func run(ctx context.Context) error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	pool, err := pgxpool.New(ctx,
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
			os.Getenv("WILSON_DB_USER"),
			os.Getenv("WILSON_DB_PASSWORD"),
			os.Getenv("WILSON_DB_HOST"),
			os.Getenv("WILSON_DB_PORT"),
			os.Getenv("WILSON_DB_DATABASE"),
		),
	)
	if err != nil {
		return err
	}
	defer pool.Close()

	if err = pool.Ping(ctx); err != nil {
		return err
	}

	r := chi.NewMux()
	r.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer)
	r.Mount("/", api.Handler())
	select {
	case <-ctx.Done():
		return nil

	}
}

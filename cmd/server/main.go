package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sviatilnik/sso/internal/sso/infrastructure/config"
	"github.com/sviatilnik/sso/internal/sso/infrastructure/interfaces/http/handlers"
)

func main() {
	conf := config.NewEnvConfig()

	db, err := sql.Open("pgx", conf.DatabaseConnectionString)
	if err != nil || db == nil {
		slog.Error("failed to open database: %v", err)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	mux.Post("/api/v1/login", handlers.NewAuthHandler().Login)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.Port),
		Handler: mux,
	}

	go func() {
		slog.Info("starting server ...")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down...")

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

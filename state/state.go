package state

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type AppState struct {
	db     *sqlx.DB
	router http.Handler
}

func New(db *sqlx.DB) *AppState {
	return &AppState{
		db:     db,
		router: loadRoutes(db),
	}
}

func (a *AppState) Start(_ context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}
	return nil
}

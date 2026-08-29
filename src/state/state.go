package state

import (
	"context"
	"fmt"
	"net/http"
)

type AppState struct {
	router http.Handler
}

func New() *AppState {
	return &AppState{
		router: loadRoutes(),
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

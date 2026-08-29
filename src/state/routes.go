package state

import (
	"Bank/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func loadRoutes(db *sqlx.DB) *chi.Mux {
	router := chi.NewRouter()
	expensesRouter := &handler.Expenses{
		Db: db,
	}

	router.Use(middleware.Logger)

	router.Post("/", expensesRouter.Create)
	router.Get("/", expensesRouter.List)
	router.Get("/{id}", expensesRouter.GetById)
	router.Patch("/{id}", expensesRouter.Update)
	router.Delete("/{id}", expensesRouter.Delete)
	router.Get("/summary", expensesRouter.GetSummary)
	return router
}

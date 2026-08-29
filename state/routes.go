package state

import (
	"Bank/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func LoadRoutes(db *sqlx.DB) *chi.Mux {
	router := chi.NewRouter()
	expensesRouter := &handler.Expenses{
		Db: db,
	}

	router.Use(middleware.Logger)

	// only for testing purposes
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	router.Post("/expenses", expensesRouter.Create)
	router.Get("/expenses", expensesRouter.List)
	router.Get("/expenses/{id}", expensesRouter.GetById)
	router.Patch("/expenses/{id}", expensesRouter.Update)
	router.Delete("/expenses/{id}", expensesRouter.Delete)
	router.Get("/expenses/summary", expensesRouter.GetSummary)
	return router
}

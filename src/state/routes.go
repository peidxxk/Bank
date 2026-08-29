package state

import (
	"Bank/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/expenses", loadExpensesRoutes)
	return router
}

func loadExpensesRoutes(router chi.Router) {
	expensesRouter := &handler.Expenses{}

	router.Post("/", expensesRouter.Create)
	router.Get("/", expensesRouter.List)
	router.Get("/{id}", expensesRouter.GetById)
	router.Patch("/{id}", expensesRouter.Update)
	router.Delete("/{id}", expensesRouter.Delete)
	router.Get("/summary", expensesRouter.GetSummary)
}

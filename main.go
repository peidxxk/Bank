package main

import (
	"Bank/state"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	_ "Bank/docs"
)

// @title BankApi
// @version 1.0
// @description API for managing expenses.
// @host localhost:3000
// @BasePath /
func main() {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	go fmt.Println("==== BankApi (TakeHome) Started ====")

	app := state.New(db)
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}

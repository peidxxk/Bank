package main

import (
	"Bank/state"
	"context"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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

	app := state.New(db)
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}

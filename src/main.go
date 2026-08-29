package main

import (
	"Bank/state"
	"context"
)

func main() {
	app := state.New()

	err := app.Start(context.TODO())
	if err != nil {
		panic(err)
	}
}

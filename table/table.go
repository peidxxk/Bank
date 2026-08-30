package table

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func CreateExpenses(db *sqlx.DB, ctx context.Context) (sql.Result, error) {
	return db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS expenses (
		id UUID PRIMARY KEY DEFAULT uuidv4(),
		amount DECIMAL NOT NULL CHECK (amount > 0),
		category TEXT NOT NULL,
		note TEXT,
		spent_on DATE NOT NULL,
		created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
}

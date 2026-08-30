package test

import (
	"Bank/state"
	"Bank/table"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"uuid"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func cleanExpenses(t *testing.T, db *sqlx.DB) {
	t.Helper()

	_, err := db.Exec(`TRUNCATE TABLE expenses`)
	if err != nil {
		t.Fatalf("clean expenses: %v", err)
	}
}

func setupTestDB(t *testing.T) *sqlx.DB {
	t.Helper()
	ctx := context.Background()

	db, err := sqlx.ConnectContext(ctx, "postgres", os.Getenv("DATABASE_TEST_URL"))

	if err != nil {
		t.Fatalf("connect to database: %v", err)
	}
	t.Cleanup(func() {
		err := db.Close()
		if err != nil {
			return
		}
	})

	_, err = table.CreateExpenses(db, ctx)
	if err != nil {
		_ = fmt.Errorf("error while creating table 'expenses': %v", err)
	}

	return db
}

func setupRouter(t *testing.T) (*sqlx.DB, http.Handler) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("load .env: %v", err)
	}

	t.Helper()
	db := setupTestDB(t)
	router := state.LoadRoutes(db)
	return db, router
}

func createTestExpense(t *testing.T, db *sqlx.DB) uuid.UUID {
	t.Helper()

	id := uuid.New()

	_, err := db.Exec(
		` INSERT INTO expenses ( id, amount, category, note, spent_on ) VALUES ($1, $2, $3, $4, $5) `,
		id, 100, "food", "Test expense", "2026-08-29")

	if err != nil {
		t.Fatalf("create test expense: %v", err)
	}

	return id
}

func TestCreateExpense(t *testing.T) {
	db, router := setupRouter(t)

	req := httptest.NewRequest(
		http.MethodPost,
		"/expenses",
		strings.NewReader(`{
			"amount": 100,
			"category": "food",
			"note": "Lunch",
			"spent_on": "2026-08-29"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d: %s",
			http.StatusCreated,
			rec.Code,
			rec.Body.String(),
		)
	}

	var count int

	err := db.Get(&count, `
		SELECT COUNT(*)
		FROM expenses
		WHERE amount = 100
		  AND category = 'food'
		  AND note = 'Lunch'
		  AND spent_on = '2026-08-29'
	`)
	if err != nil {
		t.Fatalf("check created expense: %v", err)
	}

	if count != 1 {
		t.Fatalf("expected 1 created expense, got %d", count)
	}

	cleanExpenses(t, db)
}

func TestListExpenses(t *testing.T) {
	db, router := setupRouter(t)

	createTestExpense(t, db)

	req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var expenses []struct {
		ID uuid.UUID `db:"id"`
	}

	if err := json.NewDecoder(rec.Body).Decode(&expenses); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if len(expenses) == 0 {
		t.Fatal("no expenses")
	}

	cleanExpenses(t, db)
}

func TestGetExpenseByID(t *testing.T) {
	db, router := setupRouter(t)

	id := createTestExpense(t, db)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/expenses/%s", id), nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var expense struct {
		ID uuid.UUID `db:"id"`
	}

	if err := json.NewDecoder(rec.Body).Decode(&expense); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if expense.ID != id {
		t.Fatalf("expected id %s, got %s", id, expense.ID)
	}

	cleanExpenses(t, db)
}

func TestUpdateExpenseByID(t *testing.T) {
	db, router := setupRouter(t)
	id := createTestExpense(t, db)
	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/expenses/%s", id),
		strings.NewReader(`{ "amount": 100, "category": "tech", "note": "Updated" }`))

	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var amount float64
	var category string
	var note string

	err := db.QueryRow(` SELECT amount, category, note FROM expenses WHERE id = $1 `, id).Scan(&amount, &category, &note)

	if err != nil {
		t.Fatalf("check expense: %v", err)
	}
	if amount != 100 {
		t.Fatalf("expected 100, got %f", amount)
	}

	if category != "tech" {
		t.Fatalf("expected category 'tech', got %s", category)
	}

	if note != "Updated" {
		t.Fatalf("expected note 'Updated', got %s", note)
	}
	cleanExpenses(t, db)
}

func TestDeleteExpenseByID(t *testing.T) {
	db, router := setupRouter(t)

	id := createTestExpense(t, db)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/expenses/%s", id), nil)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK && rec.Code != http.StatusNoContent {
		t.Fatalf("expected 200 OK, got %d", rec.Code)
	}

	var count int

	err := db.Get(&count, `SELECT COUNT(*) FROM expenses WHERE id = $1`, id)

	if err != nil {
		t.Fatalf("check expense: %v", err)
	}

	if count != 0 {
		t.Fatalf("expected 0 expense, got %d", count)
	}

	cleanExpenses(t, db)

}

func TestGetSummary(t *testing.T) {
	db, router := setupRouter(t)

	_ = createTestExpense(t, db)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/expenses/summary?category=food"), nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var total float64

	if err := json.NewDecoder(rec.Body).Decode(&total); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if total != 100 {
		t.Fatalf("expected 100, got %f", total)
	}

	cleanExpenses(t, db)
}

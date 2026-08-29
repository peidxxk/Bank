package handler

import (
	"Bank/dto/expense"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Expenses struct {
	Db *sqlx.DB
}

// Create
// @Summary Create expense
// @Description Creates a new expense
// @Tags expenses
// @Accept json
// @Produce json
// @Param expense body expense.CreateDto true "Expense"
// @Success 201
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /expenses [post]
func (e *Expenses) Create(w http.ResponseWriter, r *http.Request) {
	var dto expense.CreateDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if dto.Amount <= 0 {
		http.Error(w, "amount_must_be_positive", http.StatusBadRequest)
		return
	}

	result, err := e.Db.ExecContext(
		r.Context(),
		`
		INSERT INTO expenses (amount, category, note, spent_on)	VALUES ($1, $2, $3, $4)
		`,
		dto.Amount,
		dto.Category,
		dto.Note,
		dto.SpentOn,
	)
	if err != nil {
		log.Printf("create expense: %v", err)
		http.Error(w, "failed_to_create_expense", http.StatusInternalServerError)
		return
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Printf("create expense 2: %v", err)
		http.Error(w, "failed_to_create_expense", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// List
// @Summary List expenses
// @Description Returns all expenses
// @Tags expenses
// @Produce json
// @Success 200 {array} expense.ResponseDto
// @Failure 500 {string} string
// @Router /expenses [get]
func (e *Expenses) List(w http.ResponseWriter, r *http.Request) {
	rows, err := e.Db.QueryxContext(r.Context(), `SELECT * FROM expenses ORDER BY created_on DESC`)
	if err != nil {
		log.Printf("list expenses: %v", err)
		http.Error(w, "failed_to_list_expense", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var expenses []expense.ResponseDto

	for rows.Next() {
		var item expense.ResponseDto

		if err := rows.StructScan(&item); err != nil {
			log.Printf("list expenses: %v", err)
			http.Error(w, "failed_to_list_expense", http.StatusInternalServerError)
			return
		}

		expenses = append(expenses, item)
	}

	if err := rows.Err(); err != nil {
		log.Printf("iterate expenses: %v", err)
		http.Error(w, "failed_to_list_expenses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(expenses); err != nil {
		log.Printf("encode expenses: %v", err)
	}
}

// GetById TODO: Single entry
func (e *Expenses) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Expenses] GetById")
}

// Update TODO: Update entry
func (e *Expenses) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Expenses] Update")
}

// Delete TODO: Delete (hard) entry
func (e *Expenses) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Expenses] Delete")
}

// GetSummary TODO: Return by category
func (e *Expenses) GetSummary(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[Expenses] GetSummary, remoteAddr: %s\n", r.RemoteAddr)
}

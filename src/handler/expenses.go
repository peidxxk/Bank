package handler

import (
	"fmt"
	"net/http"
)

type Expenses struct{}

// Create TODO: New entry
func (e *Expenses) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Expenses] Create")
}

// List TODO: Newest first
func (e *Expenses) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Expenses] Get")
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

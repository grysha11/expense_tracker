package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grysha11/expense_tracker/internal/middleware"
)

type ExpenseHandler struct {

}

func (e ExpenseHandler) ListExpenses(w http.ResponceWriter, r *http.Request) {}
func (e ExpenseHandler) GetExpenses(w http.ResponceWriter, r *http.Request) {}
func (e ExpenseHandler) CreateExpenses(w http.ResponceWriter, r *http.Request) {}
func (e ExpenseHandler) UpdateExpenses(w http.ResponceWriter, r *http.Request) {}
func (e ExpenseHandler) DeleteExpenses(w http.ResponceWriter, r *http.Request) {}

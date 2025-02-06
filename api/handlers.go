package api

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type ExpenseHandler struct {

}

func listExpenses() []*Expense {
	return Expenses
}

func getExpense(id string) *Expense {
	for _, expense := range Expenses {
		if id == expense.ID {
			return expense
		}
	}
	return nil
}

func addExpense(expense Expense) {
	Expenses = append(Expenses, &expense)
}

func deleteExpense(id string) *Expense {
	for i, expense := range Expenses {
		if id == expense.ID {
			Expenses = append(Expenses[:i], Expenses[i+1:]...)
			return expense
		}
	}
	return nil
}

func storeExpense(expense Expense) {
	Expenses = append(Expenses, &expense)
}

func updateExpense(id string, change Expense) *Expense {
	for i, expense := range Expenses {
		if id == expense.ID {
			Expenses[i] = &change
			return expense
		}
	}
	return nil
}

func (e ExpenseHandler) ListExpenses(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(listExpenses())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (e ExpenseHandler) GetExpenses(w http.ResponseWriter, r *http.Request) {
	expense := getExpense(chi.URLParam(r, "id"))
	if expense == nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	err := json.NewEncoder(w).Encode(expense)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (e ExpenseHandler) CreateExpenses(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	err := json.NewDecoder(r.Body).Decode(expense)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	storeExpense(expense)
	err = json.NewEncoder(w).Encode(expense)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (e ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	updateExpense := updateExpense(id, expense)
	if updateExpense == nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(updateExpense)
	if err !=  nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (e ExpenseHandler) DeleteExpenses(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	expense := deleteExpense(id)
	if expense == nil {
		http.Error(w, "Expense not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

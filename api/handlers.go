package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type ExpenseHandler struct {
	DB *sql.DB
}

func (e ExpenseHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := e.DB.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Balance)
		if err != nil {
			http.Error(w, "Failed to scan users", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (e ExpenseHandler) ListExpenses(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	query := "SELECT * FROM expenses WHERE user_id = ?"
	rows, err := e.DB.Query(query, id)
	if err != nil {
		http.Error(w, "Failed to retrieve expenses", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var expense Expense
		err := rows.Scan(&expense.Id, &expense.UserId, &expense.Category, &expense.Amount, &expense.Date, &expense.Notes);
		if err != nil {
			http.Error(w, "Failed to scan expense", http.StatusInternalServerError)
			return
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Failed to retrieve expenses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}

func (e ExpenseHandler) CreateExpenses(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO expenses (user_id, category, amount, date, notes) VALUES (?, ?, ?, ?, ?)"
	_, err = e.DB.Exec(query, expense.UserId, expense.Category, expense.Amount, expense.Date, expense.Notes)
	if err != nil {
		http.Error(w, "Failed to create expense", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}

func (e ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE expenses SET user_id = ?, category = ?, amount = ?, date = ?, notes = ? WHERE id = ?"
	_, err = e.DB.Exec(query, expense.UserId, expense.Category, expense.Amount, expense.Date, expense.Notes, id)
	if err != nil {
		http.Error(w, "Failed to update expense", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expense)
}

func (e ExpenseHandler) DeleteExpenses(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	query := "DELETE FROM expenses WHERE id = ?"
	res, err := e.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
		return
	}

	checkDelete, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to retrieve affected rows", http.StatusInternalServerError)
		return
	}

	if checkDelete == 0 {
		http.Error(w, "Expense not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

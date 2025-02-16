package api

import (
	"strconv"
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"time"
)

type ExpenseHandler struct {
	DB *sql.DB
}

func checkExpense(expense Expense, userId int) bool {
	if expense.Id <= 0 || userId <= 0 || expense.Category == "" ||
		expense.Amount <= 0 || expense.Notes == "" {
			return false
		}
		return true
}

func checkUser(user User) bool {
	if user.UserId <= 0 || user.Name == "" || user.Balance <= 0 {
			return false
		}
		return true
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
		err := rows.Scan(&user.UserId, &user.Name, &user.Balance)
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
	strUserId := chi.URLParam(r, "user_id")
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		http.Error(w, "Failed to convert user_id", http.StatusInternalServerError)
		return
	}
	query := "SELECT * FROM expenses WHERE user_id = ?"
	rows, err := e.DB.Query(query, userId)
	if err != nil {
		http.Error(w, "Failed to retrieve expenses", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var expense Expense
		var dateStr string
		err := rows.Scan(&expense.Id, &expense.UserId, &expense.Category, &expense.Amount, &dateStr, &expense.Notes);
		if err != nil {
			http.Error(w, "Failed to scan expense: "+err.Error(), http.StatusInternalServerError)
			return
		}
		expense.Date, err = time.Parse("2006-01-02 15:04:05", dateStr)
		if err != nil {
			http.Error(w, "Failed to convert date: "+err.Error(), http.StatusInternalServerError)
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

func (e ExpenseHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if !checkUser(user) {
		http.Error(w, "Invalid user values", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO users (name, balance) VALUES (?, ?)"
	_, err = e.DB.Exec(query, user.Name, user.Balance)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (e ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	strUserId := chi.URLParam(r, "user_id")
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		http.Error(w, "Failed to convert user_id", http.StatusInternalServerError)
		return
	}

	if !checkExpense(expense, userId) {
		http.Error(w, "Invalid expense values", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO expenses (user_id, category, amount, notes) VALUES (?, ?, ?, ?)"
	_, err = e.DB.Exec(query, userId, expense.Category, expense.Amount, expense.Notes)
	if err != nil {
		http.Error(w, "Failed to create expense", http.StatusInternalServerError)
		return
	}

	expense.UserId = userId
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

	dateStr := expense.Date.Format("2006-01-02 15:04:05")
	expense.Date, err = time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		http.Error(w, "Failed to convert time: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if checkExpense(expense, expense.UserId) {
		http.Error(w, "Invalid expense values", http.StatusBadRequest)
		return
	}

	query := "UPDATE expenses SET user_id = ?, category = ?, amount = ?, date = ?, notes = ? WHERE id = ?"
	_, err = e.DB.Exec(query, expense.UserId, expense.Category, expense.Amount, expense.Date, expense.Notes, id)
	if err != nil {
		http.Error(w, "Failed to update expense: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expense)
}

func (e ExpenseHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	user_id := chi.URLParam(r, "user_id")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if checkUser(user) {
		http.Error(w, "Invalid user values", http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET user_id = ?, name = ?, balance = ? WHERE user_id = ?"
	_, err = e.DB.Exec(query, user.UserId, user.Name, user.Balance, user_id)
	if err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (e ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
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

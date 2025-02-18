package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grysha11/expense_tracker/api"
	"github.com/grysha11/expense_tracker/db"
)

const SERVER_PORT string = ":8080"

func ExpenseRoutes(db *sql.DB) chi.Router {
	expenseHandler := api.ExpenseHandler{DB: db}
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", expenseHandler.ListUsers)
	r.Post("/", expenseHandler.CreateUser)
	r.Put("/{user_id}", expenseHandler.UpdateUser)
	r.Post("/{user_id}/expenses", expenseHandler.CreateExpense)
	r.Get("/{user_id}/expenses", expenseHandler.ListExpenses)
	r.Put("/{user_id}/expenses/{id}", expenseHandler.UpdateExpense)
	r.Delete("/{user_id}/expenses/{id}", expenseHandler.DeleteExpense)

	return r
}

func main() {
	database, err := sql.Open("mysql", db.GetDSN())
	if err != nil {
		fmt.Println("error initializing db")
		panic(err.Error())
	}
	defer database.Close()

	err = database.Ping()
	if err != nil {
		fmt.Println("error ping db")
		panic(err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(api.ErrorHandler)

	fmt.Println("starting api and listening on port" + SERVER_PORT + "...")

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Mount("/users", ExpenseRoutes(database))
	err = http.ListenAndServe(SERVER_PORT, r)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

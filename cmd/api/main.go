package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grysha11/expense_tracker/api"
)

func ExpenseRoutes() chi.Router {
	r := chi.NewRouter()
	expenseHandler := api.ExpenseHandler{}
	r.Get("/", expenseHandler.ListExpenses)
	r.Post("/", expenseHandler.CreateExpenses)
	r.Get("/{id}", expenseHandler.GetExpenses)
	r.Put("/{id}", expenseHandler.UpdateExpense)
	r.Delete("/{id}", expenseHandler.DeleteExpenses)

	return r
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fmt.Println("starting api")

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Mount("/expenses", ExpenseRoutes())
	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

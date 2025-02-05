package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/grysha11/expense_tracker/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fmt.Println("starting api")

	r.Get("/", func(w http.ResponceWriter, r *http.Request)) {
		w.write([]byte("OK"))
	}
	r.Mount("/expenses", ExpenseRoutes())
	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		w.write([]byte("Error"))
	}
}

func ExpenseRoutes() chi.Router {
	r := chi.NewRouter()
	expenseHandler := ExpenseHandler{}
	r.Get("/", expenseHandler.ListExpenses)
	r.Post("/", expenseHandler.CreateExpenses)
	r.Get("/{id}", expenseHandler.GetExpenses)
	r.Put("/{id}", expenseHandler.UpdateExpenses)
	r.Delete("/{id}", expenseHandler.DeleteExpenses)

	return r
}

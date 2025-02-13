package api

import "time"

type User struct {
	UserId int `json:"user_id"`
	Name string `json:"name"`
	Balance float32 `json:"balance"`
}

type Expense struct {
	Id int `json:"id"`
	UserId string `json:"user_id"`
	Category string `json:"category"`
	Amount float32 `json:"amount"`
	Date time.Time `json:"date"`
	Notes string `json:"notes"`
}
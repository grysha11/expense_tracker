package api

type User struct {
	UserId int `json:"user_id"`
	Name string `json:"name"`
	Balance float32 `json:"balance"`
}

type Expense struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Category string `json:"category"`
	Amount float32 `json:"amount"`
	Date string `json:"date"`
	Notes string `json:"notes"`
}
package api

type Expense struct {
	ID			string `json:"id"`
	UserId		string `json:"user_id"`
	Category	string `json:"category"`
	Amount		float32 `json:"amount"`
	Currency	string `json:"currency"`
	Date		string `json:"date"`
	Notes		string `json:"notes"`
}

var Expenses = []*Expense {
	{
		ID:			"123",
		UserId:		"u123",
		Category:	"grocery",
		Amount:		46.42,
		Currency:	"USD",
		Date:		"2024-02-04T12:30:00Z",
		Notes:		"food shopping",
	},
}

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
	Salary  float32 `json:"salary"`
	Tax     float32 `json:"tax"`
}

var test_user = []user{
	{Name: "Admin", Balance: 5052.1, Salary: 7500.24, Tax: 27.0},
	{Name: "Admin2", Balance: 0.0, Salary: 500.99, Tax: 15.0},
}

func getTest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, test_user)
}

func runTest() {
	router := gin.Default()
	router.GET("/test_user", getTest)
	router.Run("localhost:8000")
}

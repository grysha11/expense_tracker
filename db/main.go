package db

import "os"

func GetDSN() string {
	dbHost := os.Getenv("DB_HOST_KEY")
	dbPort := os.Getenv("DB_PORT_KEY")
	dbUser := os.Getenv("DB_USER_KEY")
	dbPass := os.Getenv("DB_PASS_KEY")
	dbName := os.Getenv("DB_NAME_KEY")

	return dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}

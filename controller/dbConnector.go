package controller

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"pp-api/config"

	//this is needed
	_ "github.com/go-sql-driver/mysql"
)

var connection *sqlx.DB

// Connect to the database
func connectToDatabase() error {
	var databaseConnectionErr error
	connection, databaseConnectionErr = sqlx.Connect("mysql", config.GetConfig().DbConnectionString)
	if databaseConnectionErr != nil {
		fmt.Println("connectToDatabase")
		fmt.Println(databaseConnectionErr.Error())
		return databaseConnectionErr
	}

	return nil
}

//GetConnection Returns connection to MySQL
func GetConnection() *sqlx.DB {
	if connection == nil {
		connectToDatabase()
	}
	return connection
}

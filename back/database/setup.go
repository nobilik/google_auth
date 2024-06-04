package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {

	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	// Формирование строки подключения
	dsn := fmt.Sprintf("%s:%s@tcp(dbContainer)/%s", dbUser, dbPassword, dbName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	time.Sleep(10 * time.Second)
	// test the connection
	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}
	createUsersTable()
	createGoogleProfilesTable()
	fmt.Println("Connected to MySQL database!")
}

func createUsersTable() {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			created_at DATETIME,
			updated_at DATETIME,
			full_name VARCHAR(255),
			email VARCHAR(100) NOT NULL UNIQUE,
			telephone VARCHAR(20),
			password VARCHAR(255)
		);
	`

	_, err := DB.Exec(createTableQuery)
	if err != nil {
		panic(err.Error())
	}
}

func createGoogleProfilesTable() {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS google_profiles (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL UNIQUE,
			created_at DATETIME,
			updated_at DATETIME,
			uid VARCHAR(255) NOT NULL UNIQUE
		);
	`

	_, err := DB.Exec(createTableQuery)
	if err != nil {
		panic(err.Error())
	}
}

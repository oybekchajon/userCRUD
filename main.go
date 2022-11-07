package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	PostgresUser     = "postgres"
	PostgresPassword = "1234"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresDatabase = "postgres"
)

func main() {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPassword,
		PostgresDatabase)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connection database: %v", err)
	}

	dbManager := NewDBManager(db)

	// var createdUser User
	// createdUser.Name = "Hasan"
	// createdUser.Age = 24

	// dbManager.CreateUser(&createdUser)

	// var updatedUser User

	// updatedUser.Id = 6
	// updatedUser.Name = "Green"
	// updatedUser.Age = 55

	// dbManager.UpdateUser(&updatedUser)


	userinfo, err := dbManager.Get(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userinfo)	
}
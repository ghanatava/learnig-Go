package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	createTable(db)
	insertData(db)
	selectData(db)
	updateData(db)
	deleteData(db)

}

func createTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50),
			age INT
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("CREATE TABLE")
}

func insertData(db *sql.DB) {
	query := `
	INSERT INTO users (name, age)
	VALUES ('Alice', 25), ('Bob', 30)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert users")
}

func updateData(db *sql.DB) {
	query := `UPDATE users SET age = $1 WHERE name = $2`
	_, err := db.Exec(query, 35, "Alice")
	if err != nil {
		panic(err)
	}
	fmt.Println("Data updated successfully")
}

func deleteData(db *sql.DB) {
	query := `DELETE FROM users WHERE name = $1`
	_, err := db.Exec(query, "Bob")
	if err != nil {
		panic(err)
	}
	fmt.Println("Data deleted successfully")
}

func selectData(db *sql.DB) {
	query := `
		SELECT * FROM users;
	`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

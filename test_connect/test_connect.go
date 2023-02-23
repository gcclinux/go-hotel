package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Created to Lesson 98. How to connect a Go application to a database

func main() {
	// connect to a database

	conn, err := sql.Open("pgx", "host=odroid port=5432 dbname=hotel_test user=testie password=pastie")
	if err != nil {
		log.Fatal(fmt.Sprintf("Uanble to connect %v\n", err))
	}
	defer conn.Close()
	log.Println("Connected to database!")

	// test my connection
	err = conn.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot ping database %v\n", err))
	}
	log.Println("Pinged database!")
	fmt.Println("---------------------------------------")

	// get rows from table

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot get getAllRows %v\n", err))
	}

	// insert to row

	query := `insert into users(first_name, last_name) values ($1, $2)`
	_, err = conn.Exec(query, "Jack", "Brown")
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot insert to users %v\n", err))
	}

	log.Println("Inserted a row!")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot get getAllRows %v\n", err))
	}

	// update a row
	stmt := `update users set first_name = $1 where id = $2`
	_, err = conn.Exec(stmt, "Jackie", 5)

	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot update users %v\n", err))
	}

	log.Println("Updated a row!")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot get getAllRows %v\n", err))
	}

	// get one row by id
	query = `Select id, first_name, last_name from users where id = $1`
	var firstName, lastName string
	var id int
	row := conn.QueryRow(query, 1)

	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Selected id: ", id, firstName, lastName)

	// delete a row
	query = `delete from users where id = $1`
	_, err = conn.Exec(query, 6)
	if err != nil {
		log.Println(err)
	}
	log.Println("Deleted a row")

	// get rows again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot get getAllRows %v\n", err))
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("Select id, first_name, last_name from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is:", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(fmt.Sprintf("Error scanning rows %v\n", err))
	}

	fmt.Println("---------------------------------------")

	return nil
}

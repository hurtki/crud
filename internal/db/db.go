package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDataBase() *sql.DB {

	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	postgres_db := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("postgres://%s:%s@postgres:5432/%s", postgres_user, postgres_password, postgres_db)

	fmt.Printf("dsn: %s\n", dsn)
	var db *sql.DB
	var err error

	// Retry loop
	for i := 0; i < 10; i++ {
		db, err = sql.Open("pgx", dsn)
		if err != nil {
			fmt.Println("failed to open db:", err)
		} else if err = db.PingContext(context.Background()); err == nil {
			break
		}

		fmt.Println("waiting for db to be ready...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		panic("cannot connect to db: " + err.Error())
	}

	if err != nil {
		fmt.Println("error occured")
		return nil
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS USERS (
		name text,
		age int,
		id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY
	);`)

	if err != nil {
		fmt.Println("cannot create table")
	}
	var id int
	err = db.QueryRow(
		`INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`,
		"alice", 12,
	).Scan(&id)
	if err != nil {
		fmt.Println("insert error:", err)
	} else {
		fmt.Println("Inserted user id:", id)
	}

	res, err := db.Query("SELECT * FROM USERS")

	if err != nil {
		fmt.Println("error occured while execing querry: " + err.Error())
	}

	for res.Next() {
		var id int
		var name string
		var age int
		if err := res.Scan(&name, &age, &id); err != nil {
			fmt.Println("error scanning")
		}
		fmt.Println(id, name, age)
	}

	return db
}

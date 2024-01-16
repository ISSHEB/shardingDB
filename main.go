package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	Name    string
	ConnStr string
	DB      *sql.DB
}

func main() {
	connStr1 := "user=postgres password=postgres dbname=db1 sslmode=disable"
	db1, err := NewDatabase("db1", connStr1)
	if err != nil {
		log.Fatal(err)
	}
	
	connStr2 := "user=postgres password=postgres dbname=db2 sslmode=disable"
	db2, err := NewDatabase("db2", connStr2)
	if err != nil {
		log.Fatal(err)
	}
	
	leastUsedDB, err := getDatabaseWithLessMemory(db1, db2)
	if err != nil {
		log.Fatal(err)
	}
	
	s := SendToDB(leastUsedDB.DB)
	
	fmt.Printf("Data %s sent successfully.", s)
}

func NewDatabase(name, connStr string) (*Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	
	// Test the connection
	err = db.Ping()
	
	if err != nil {
		return nil, err
	}
	
	defer db.Close()
	
	return &Database{Name: name, ConnStr: connStr, DB: db}, nil
}

// Функция для получения размера базы данных
func getDatabaseSize(db *sql.DB, dbName string) (int64, error) {
	var size int64
	err := db.QueryRow("SELECT pg_database_size($1)", dbName).Scan(&size)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// Функция для определения базы данных с наименьшим размером
func getDatabaseWithLessMemory(db1 *Database, db2 *Database) (*Database, error) {
	size1, err := getDatabaseSize(db1.DB, "db1")
	if err != nil {
		return nil, err
	}
	size2, err := getDatabaseSize(db2.DB, "db2")
	if err != nil {
		return nil, err
	}
	
	if size1 < size2 {
		return db1, nil
	}
	
	return db2, nil
}

// Функция для отправки данных в базу данных
func SendToDB(db *sql.DB) error {
	query := "INSERT INTO my_table (column1, column2) VALUES ($1, $2)"
	args := []interface{}{"value1", "value2"}
	
	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}
	
	return nil
}

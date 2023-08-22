package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)


type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	port := os.Getenv("DATABASE_PORT") 
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD") 
	dbname := os.Getenv("DATABASE_NAME")

	psqlInfo := fmt.Sprintf("port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
     port, user, password, dbname)
    
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

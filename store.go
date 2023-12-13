package main

import "database/sql"

type Store interface {
	// users
	CreateUser() error
}

type Storage struct {
	// some fields
	db *sql.DB
}

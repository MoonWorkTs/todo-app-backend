package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListTable  = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host       string
	Port       string
	Username   string
	Password   string
	DBName     string
	SSLMode    string
	SearchPath string
}

func NewPostresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode, cfg.SearchPath))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

package database

import (
	"database/sql"
	"fmt"
)

type ConnectionInfo struct {
	Host     string
	Port     int
	UserName string
	DBMame   string
	SSLMode  string
	Password string
}

func NewPostgresConnection(info ConnectionInfo) *sql.DB {
	db, _ := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		info.Host, info.Port, info.UserName, info.DBMame, info.SSLMode, info.Password))
	return db
}

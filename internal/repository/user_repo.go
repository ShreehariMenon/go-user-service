package repository

import (
    "database/sql"
    "go-user-service/db/sqlc"
)

type Repo struct {
    Q *db.Queries
}

func NewRepo(dbConn *sql.DB) *Repo {
    return &Repo{Q: db.New(dbConn)}
}
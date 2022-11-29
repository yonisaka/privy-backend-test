package repositories

import "github.com/jmoiron/sqlx"

func NewBaseRepo(db *sqlx.DB) *BaseRepo {
	return &BaseRepo{db: db}
}

type BaseRepo struct {
	db *sqlx.DB
}

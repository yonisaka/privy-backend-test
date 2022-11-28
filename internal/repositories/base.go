package repositories

import "gorm.io/gorm"

func NewBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{db: db}
}

type BaseRepo struct {
	db *gorm.DB
}

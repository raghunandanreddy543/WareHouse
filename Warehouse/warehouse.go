package Warehouse

import (
	"context"
	"database/sql"
)

type UserRepositoryInterface interface {
	GetQuantity(ctx context.Context) error
}

type userRepository struct {
	db *sql.DB
}

func (u userRepository) GetQuantity(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &userRepository{
		db,
	}
}

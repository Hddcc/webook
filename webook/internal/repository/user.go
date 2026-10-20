package repository

import (
	"context"
	"webook/internal/domain"
)

type UserRepository struct {
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	panic("UserRepository.Create is not implemented")
}

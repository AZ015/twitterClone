package domain

import (
	"context"
	"twitter"
	"twitter/uuid"
)

type UserService struct {
	UserRepo twitter.UserRepo
}

func NewUserService(ur twitter.UserRepo) *UserService {
	return &UserService{
		UserRepo: ur,
	}
}

func (us *UserService) GetByID(ctx context.Context, id string) (twitter.User, error) {
	if !uuid.Validate(id) {
		return twitter.User{}, twitter.ErrInvalidUUID
	}
	return us.UserRepo.GetByID(ctx, id)
}

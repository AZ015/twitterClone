package graph

import (
	"context"
	"twitter"
)

func mapUser(u twitter.User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

func (q *queryResolver) Me(ctx context.Context) (*User, error) {
	userID, err := twitter.GetUserIdFromContext(ctx)
	if err != nil {
		return nil, twitter.ErrUnauthenticated
	}

	return mapUser(twitter.User{
		ID: userID,
	}), nil
}

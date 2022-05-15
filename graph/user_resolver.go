package graph

import (
	"context"
	graphql_twitter "github.com/shawnzxx/graphql-twitter"
)

// mapUser: map domain user into graph user
func mapUser(u graphql_twitter.User) *User {
	return &User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Password,
		CreatedAt: u.CreatedAt,
	}
}

func (q *queryResolver) Me(ctx context.Context) (*User, error) {
	panic("implement me")
}

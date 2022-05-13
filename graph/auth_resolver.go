package graph

import (
	"context"
	graphql_twitter "github.com/shawnzxx/graphql-twitter"
)

//mapAuthResponse: map domain AuthResponse into graph AuthResponse
func mapAuthResponse(a graphql_twitter.AuthResponse) *AuthResponse {
	return &AuthResponse{
		AccessToken: a.AccessToken,
		User:        mapUser(a.User),
	}
}

func (m *mutationResolver) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {
	res, err := m.AuthService.Register(ctx, graphql_twitter.RegisterInput{
		Email:           input.Email,
		Username:        input.Username,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
	})
	if err != nil {

	}
	return mapAuthResponse(res), nil
}
func (m *mutationResolver) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	panic("implement me")
}

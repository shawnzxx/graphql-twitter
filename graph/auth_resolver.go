package graph

import (
	"context"
	"errors"
	graphql_twitter "github.com/shawnzxx/graphql-twitter"
)

// mapAuthResponse: map domain AuthResponse into graph AuthResponse
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
		switch {
		case errors.Is(err, graphql_twitter.ErrValidation) ||
			errors.Is(err, graphql_twitter.ErrEmailTaken) ||
			errors.Is(err, graphql_twitter.ErrUsernameTaken):
			return nil, buildBadRequestError(ctx, err)
		default:
			return nil, err
		}
	}

	return mapAuthResponse(res), nil
}
func (m *mutationResolver) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	res, err := m.AuthService.Login(ctx, graphql_twitter.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		switch {
		case errors.Is(err, graphql_twitter.ErrValidation) ||
			errors.Is(err, graphql_twitter.ErrBadCredentials):
			return nil, buildBadRequestError(ctx, err)
		default:
			return nil, err
		}
	}

	return mapAuthResponse(res), nil
}

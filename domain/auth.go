package domain

import (
	"context"
	"errors"
	"fmt"
	graphql_twitter "github.com/shawnzxx/graphql-twitter"
	"golang.org/x/crypto/bcrypt"
)

var passwordCost = bcrypt.DefaultCost

type AuthService struct {
	UserRepo graphql_twitter.UserRepo
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) Register(ctx context.Context, input graphql_twitter.RegisterInput) (graphql_twitter.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return graphql_twitter.AuthResponse{}, err
	}

	//check if username is already taken
	if _, err := as.UserRepo.GetByUsername(ctx, input.Username); !errors.Is(err, graphql_twitter.ErrNotFound) {
		return graphql_twitter.AuthResponse{}, graphql_twitter.ErrUsernameTaken
	}

	//check if email is already taken
	if _, err := as.UserRepo.GetByEmail(ctx, input.Email); !errors.Is(err, graphql_twitter.ErrNotFound) {
		return graphql_twitter.AuthResponse{}, graphql_twitter.ErrEmailTaken
	}

	//initial new user
	user := graphql_twitter.User{
		Username: input.Username,
		Email:    input.Email,
	}

	//hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), passwordCost)
	if err != nil {
		return graphql_twitter.AuthResponse{}, fmt.Errorf("error hashing password %v", err)
	}
	user.Password = string(hashPassword)

	//create the user
	user, err = as.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return graphql_twitter.AuthResponse{}, fmt.Errorf("error creating user: %v", err)
	}

	//return accessToken and user
	return graphql_twitter.AuthResponse{
		AccessToken: "JWT",
		User:        user,
	}, nil
}

func (as *AuthService) Login(ctx context.Context, input graphql_twitter.LoginInput) (graphql_twitter.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return graphql_twitter.AuthResponse{}, err
	}

	user, err := as.UserRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, graphql_twitter.ErrNotFound):
			return graphql_twitter.AuthResponse{}, graphql_twitter.ErrBadCredentials
		default:
			return graphql_twitter.AuthResponse{}, err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return graphql_twitter.AuthResponse{}, graphql_twitter.ErrBadCredentials
	}

	return graphql_twitter.AuthResponse{
		AccessToken: "JWT",
		User:        user,
	}, nil
}

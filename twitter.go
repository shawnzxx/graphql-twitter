package graphql_twitter

import "errors"

var (
	ErrBadCredentials = errors.New("email or password wrong combination")
	ErrNotFound       = errors.New("not found")
	ErrValidation     = errors.New("validation error")
)

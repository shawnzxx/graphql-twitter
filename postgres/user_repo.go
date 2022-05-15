package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	graphql_twitter "github.com/shawnzxx/graphql-twitter"
)

type UserRepo struct {
	DB *DB
}

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (ur UserRepo) Create(ctx context.Context, user graphql_twitter.User) (graphql_twitter.User, error) {
	tx, err := ur.DB.Pool.Begin(ctx)
	if err != nil {
		return graphql_twitter.User{}, fmt.Errorf("error starting transaction %v", err)
	}
	defer tx.Rollback(ctx)

	user, err = createUser(ctx, tx, user)
	if err != nil {
		return graphql_twitter.User{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return graphql_twitter.User{}, fmt.Errorf("error commiting: %v", err)
	}

	return user, err
}

func createUser(ctx context.Context, tx pgx.Tx, user graphql_twitter.User) (graphql_twitter.User, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *;`

	u := graphql_twitter.User{}

	if err := pgxscan.Get(ctx, tx, &u, query, user.Username, user.Email, user.Password); err != nil {
		return graphql_twitter.User{}, fmt.Errorf("error insert: %v", err)
	}

	return u, nil
}

func (ur UserRepo) GetByUsername(ctx context.Context, username string) (graphql_twitter.User, error) {
	query := `SELECT * FROM users WHERE username = $1 LIMIT 1;`

	u := graphql_twitter.User{}

	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query, username); err != nil {
		if pgxscan.NotFound(err) {
			return graphql_twitter.User{}, graphql_twitter.ErrNotFound
		}
		return graphql_twitter.User{}, fmt.Errorf("error select query by username: %v", err)
	}
	return u, nil
}

func (ur UserRepo) GetByEmail(ctx context.Context, email string) (graphql_twitter.User, error) {
	query := `SELECT * FROM users WHERE email = $1 LIMIT 1;`

	u := graphql_twitter.User{}

	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query, email); err != nil {
		if pgxscan.NotFound(err) {
			return graphql_twitter.User{}, graphql_twitter.ErrNotFound
		}
		return graphql_twitter.User{}, fmt.Errorf("error select query by username: %v", err)
	}
	return u, nil
}

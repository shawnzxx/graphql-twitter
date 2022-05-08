package postgres

import (
	"context"
	"fmt"
	"path"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shawnzxx/graphql-twitter/config"
	"log"
)

type DB struct {
	Pool *pgxpool.Pool
	conf *config.Config
}

//New init a db poll and return db instance.
//db have functions like Ping, Close, Migrate
func New(ctx context.Context, conf *config.Config) *DB {
	dbConf, err := pgxpool.ParseConfig(conf.Database.URL)
	if err != nil {
		log.Fatalf("can't parse postgres connection string: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, dbConf)
	if err != nil {
		log.Fatalf("can't connect to postgres %v", err)
	}

	db := &DB{Pool: pool, conf: conf}

	db.Ping(ctx)

	return db
}

func (db *DB) Ping(ctx context.Context) {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Fatalf("can't ping postgres: %v", err)
	}
}

func (db *DB) Close() {
	db.Pool.Close()
}

//Migrate do migration process
func (db *DB) Migrate() error {
	//find current file location
	_, b, _, _ := runtime.Caller(0)

	migrationPath := fmt.Sprintf("file:///%s/migrations", path.Dir(b))

	m, err := migrate.New(migrationPath, db.conf.Database.URL)
	if err != nil {
		return fmt.Errorf("error create the migrate instance: %v", err)
	}

	//we catch err is not nil and err is not equal to ErrNoChange error
	//since we will do migration for every single deployment
	//we do not want to capture ErrNoChange err
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error migrate up: %v", err)
	}

	log.Println("migration done")

	return nil
}

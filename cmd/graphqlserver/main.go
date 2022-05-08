package main

import (
	"context"
	"fmt"
	"github.com/shawnzxx/graphql-twitter/config"
	"github.com/shawnzxx/graphql-twitter/postgres"
	"log"
)

func main() {
	ctx := context.Background()

	conf := config.New()

	db := postgres.New(ctx, conf)

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("WORKING")
}

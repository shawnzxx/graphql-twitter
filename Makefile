# Commands for running docker compose
COMPOSE := docker-compose -f docker-compose.yaml

setup: db-api
# teardown stops and removes all containers
teardown:
	$(COMPOSE) down

# db-api runs the db service for API defined in the compose file
db-api:
	$(COMPOSE) up -d db-api

#up till the latest file
migrate:
	migrate -source file://postgres/migrations \
			-database postgres://postgres:postgres@127.0.0.1:5432/twitter_clone_development?sslmode=disable up

#down one version
roolback:
	migrate -source file://postgres/migrations \
			-database postgres://postgres:postgres@127.0.0.1:5432/twitter_clone_development?sslmode=disable down

#delete all db we migrated
drop:
	migrate -source file://postgres/migrations \
			-database postgres://postgres:postgres@127.0.0.1:5432/twitter_clone_development?sslmode=disable drop

migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir postgres/migrations $$name

run:
	go run cmd/graphqlserver/main.go

generate:
	go generate ./graph
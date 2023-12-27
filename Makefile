swag:
	swag init --parseDependency --parseInternal --parseDepth 1 -md ./documentation -o ./docs

sql:
	sqlc generate 

migrate:
	migrate -verbose -path ./database/migrations -database $(MIGRATION_URL) up

migrate-down:
	migrate -verbose -path ./database/migrations -database $(MIGRATION_URL) down

migrate-create:
	migrate create -ext sql -dir ./database/migrations -seq $(name)

build: 
	go build -o ./bin/run ./main.go

run: migrate build
	./bin/run

.PHONY: swag sql migrate migrate-down migrate-create build run
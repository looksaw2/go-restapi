run:
	go run ./cmd/api/main.go
test:
	go test -v -cover ./...
migrate:
	migrate create -seq -ext=.sql -dir=./migrations create_movies_table
migrate_v2:
	migrate create -seq -ext=.sql -dir=./migrations add_movies_check_constraints
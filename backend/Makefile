postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=test -e POSTGRES_PASSWORD=testsecret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username test --owner=test test

dropdb:
	docker exec -it postgres14 dropdb test

migrateup:
	migrate -path db/migration/ -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_URL}:5432/${DB_TABLE}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_URL}:5432/${DB_TABLE}?sslmode=disable" -verbose down

new_migrate:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir db/migration -seq $$name

sqlc:
	sqlc generate

test:
	go test -v -cover ./... 

ssl:
	@read -p "Enter length: " len;\
		openssl rand -base64 $$len

.PHONY:  migrateup, migratedown, new_migrate, sqlc, ssl

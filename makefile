postgres:
		docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
		docker exec -it postgres12 createdb --username=postgres --owner=postgres simple_bank

dropdb:
		docker exec -it postgres12 dropdb simple_bank

migrateup:
		migrate -path db/migration -database "postgresql://postgres:secret@0.0.0.0:5432/simple_bank?sslmode=disable" force 17

migratedown:
		migrate -path db/migration -database "postgresql://postgres:secret@0.0.0.0:5432/simple_bank?sslmode=disable" force 17

sqlc:
		sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
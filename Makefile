network:
	docker network create fc-order

postgres:
	docker run --name fc-pg15 --network fc-order -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it fc-pg15 createdb --username=root --owner=root filip-club

dropdb:
	docker exec -it fc-pg15 dropdb filip-club

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/filip-club?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/filip-club?sslmode=disable" -verbose down

showdb:
	docker exec -it fc-pg15 psql -U root -d filip-club

sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/fosuch/fc-order/db/sqlc Store

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: network postgres createdb dropdb migrateup migratedown showdb sqlc test server

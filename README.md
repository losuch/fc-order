# fc-order

filip-club order backen API

## golang-migrate

Tool to maintain database schema migrations.
https://github.com/golang-migrate/migrate

### Install

`$ brew install golang-migrate`

### Create Migration file

`$ migrate create -ext sql -dir db/migration -seq init_schema`

### Migrate Up

`$ migrate -path db/migration -database "postgres://localhost:5432/filipclub?sslmode=disable" up`

## Database

Database schema design: https://dbdiagram.io/

### SQLC

SQLC is a code generation tool for writing SQL queries in Go. It is designed to replace many of the "ORM" style libraries and provide a much simpler and more performant interface to your database of choice. The sqlc.yaml contains the configuration for the sqlc tool.

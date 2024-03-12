package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var testConn *pgx.Conn
var testQueries *Queries

func TestMain(m *testing.M) {
    config, err := pgx.ParseConfig("postgresql://root:secret@localhost:5432/filip-club?sslmode=disable")
    if err != nil {
        log.Fatal("unable to parse config:", err)
    }

    testConn, err = pgx.ConnectConfig(context.Background(), config)
    if err != nil {
        log.Fatal("unable to connect to database:", err)
    }
    // defer testConn.Close(context.Background())

    testQueries = New(testConn)

    os.Exit(m.Run())
}
package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/losuch/fc-order/util"
)

var testConn *pgx.Conn
var testQueries *Queries

func TestMain(m *testing.M) {
    
    config, err := util.LoadConfig("../..")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

    dbCconfig, err := pgx.ParseConfig(config.DBSource)
    if err != nil {
        log.Fatal("unable to parse config:", err)
    }

    testConn, err = pgx.ConnectConfig(context.Background(), dbCconfig)
    if err != nil {
        log.Fatal("unable to connect to database:", err)
    }
    // defer testConn.Close(context.Background())

    testQueries = New(testConn)

    os.Exit(m.Run())
}
package api

import (
	db "github.com/losuch/fc-order/db/sqlc"
	"github.com/losuch/fc-order/util"

	"testing"

	"github.com/stretchr/testify/require"
)


func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		// TokenSymmetricKey:   util.RandomString(32),
		// AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

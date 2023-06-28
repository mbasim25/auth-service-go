package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	db "github.com/mbasim25/ticketing-app-microservices/db/sqlc"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func newTestServer(t *testing.T, store db.Store) *Server {
	server := NewServer(store)

	return server
}

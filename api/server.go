package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	db "github.com/mbasim25/ticketing-app-microservices/db/sqlc"
	"github.com/mbasim25/ticketing-app-microservices/token"
	"github.com/mbasim25/ticketing-app-microservices/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("couldn't create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	router.POST("/", server.createUser)
	router.POST("/login", server.loginUser)
	router.GET("/", server.listUsers)
	server.router = router

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

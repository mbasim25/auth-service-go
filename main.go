package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/mbasim25/void/api"
	db "github.com/mbasim25/void/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:password@localhost:5435/void?sslmode=disable"
	serverAddress = "0.0.0.0:8081"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("could not connect to the db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("could not start server", err)
	}
}

package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/mbasim25/ticketing-app-microservices/api"
	db "github.com/mbasim25/ticketing-app-microservices/db/sqlc"
	"github.com/mbasim25/ticketing-app-microservices/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load env variables", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("could not connect to the db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server", err)
	}
}

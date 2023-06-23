package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/mbasim25/void/api"
	db "github.com/mbasim25/void/db/sqlc"
	"github.com/mbasim25/void/util"
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

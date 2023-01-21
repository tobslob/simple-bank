package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tobslob/simple-bank/api"
	db "github.com/tobslob/simple-bank/db/sqlc"
	"github.com/tobslob/simple-bank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("Cannot start server:", err)
	}
}

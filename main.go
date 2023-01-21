package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tobslob/simple-bank/api"
	db "github.com/tobslob/simple-bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:secret@0.0.0.0:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err = server.Start(serverAddress); err != nil {
		log.Fatal("Cannot start server:", err)
	}
}

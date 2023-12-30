package main

import (
	"database/sql"
	"log"

	"github.com/akagami-harsh/SimpleBank/api"
	db "github.com/akagami-harsh/SimpleBank/db/sqlc"
	_ "github.com/lib/pq" // import postgres driver
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

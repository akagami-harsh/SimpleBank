package main

import (
	"database/sql"
	"log"

	"github.com/akagami-harsh/SimpleBank/api"
	db "github.com/akagami-harsh/SimpleBank/db/sqlc"
	"github.com/akagami-harsh/SimpleBank/util"
	_ "github.com/lib/pq" // import postgres driver
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

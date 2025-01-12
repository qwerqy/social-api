package main

import (
	"log"

	"github.com/qwerqy/social-api-go/internal/db"
	"github.com/qwerqy/social-api-go/internal/env"
	"github.com/qwerqy/social-api-go/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")

	conn, err := db.New(addr, 3, 3, "15m")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}

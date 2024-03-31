package main

import (
	"log"

	"github.com/ppp3ppj/wywy/config"
	"github.com/ppp3ppj/wywy/db"
	"github.com/ppp3ppj/wywy/internal/server"
)

func main() {
    conf := config.ConfigGetting()
    db := db.NewPostgresDatabase(conf.Database)
    defer func() {
        if err := db.Close(); err != nil {
            log.Fatalf("Failed to close database connection: %v", err)
        }
    }()

    server := server.NewEchoServer(conf, db)
    server.Start()
}

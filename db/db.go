package db

import (
	"log"
	"github.com/jmoiron/sqlx"
	"github.com/ppp3ppj/wywy/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func DbConnect(cfg config.IDbconfig) *sqlx.DB {
    db, err := sqlx.Connect("pgx", cfg.Url())
    if err != nil {
        log.Fatalf("connect to db failed: %v", err)
    }
    db.DB.SetMaxOpenConns(cfg.MaxConnections())
    return db
}

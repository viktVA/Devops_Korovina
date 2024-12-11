package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"test/internal/cfg"
	"test/internal/server"
)

func main() {
	cfg, err := cfg.NewCfg()
	if err != nil {
		panic(err)
	}
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	server := server.NewServer(cfg, db)
	server.Start()

}

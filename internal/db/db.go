package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alaminsobuj/goLang/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

// package-level exported DB
var DB *sql.DB

func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database Connection Error: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("DB Ping Failed: ", err)
	}

	log.Println("✅ Database Connected")
}

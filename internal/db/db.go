package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alaminsobuj8/goLang/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database Connection Error: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB Ping Failed: ", err)
	}

	log.Println("✅ Database Connected")

	return db
}

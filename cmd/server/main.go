package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alaminsobuj/goLang/internal/config"
	"github.com/alaminsobuj/goLang/internal/db"
	appHttp "github.com/alaminsobuj/goLang/internal/http"
	"github.com/alaminsobuj/goLang/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	logg := logger.New()
	// Connect to database (global db.DB will be initialized)
	db.Connect(cfg)
	// database := db.Connect(cfg)
	// _ = database // you can use it later
 
	r := http.NewServeMux()
	// r = http.NewServeMux()
	r.Handle("/", http.NewServeMux())

	router := appHttp.NewRouter()

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	go func() {
		logg.Printf("🚀 Server running on port %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🔻 Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Forced Shutdown:", err)
	}
	log.Println("✅ Server exited successfully")
}

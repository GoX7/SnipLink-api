package main

import (
	"api/internal/config"
	"api/internal/controlers"
	"api/internal/logger"
	"api/internal/sqlite"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error load config:", err)
	}

	logs, err := logger.New(cfg)
	if err != nil {
		log.Fatal("Error load logger:", err)
	}

	db, err := sqlite.New(cfg, logs)
	if err != nil {
		logs.Server.Warn("Stop server, error connect to database")
		logs.Sqlite.Error(fmt.Sprintf("Error connect database: %v", err))
		log.Fatal("Error")
	}

	logs.Sqlite.Info("Connect sqlite")
	logs.Server.Info("Starting (1/2)...")
	start(cfg, logs, db)
}

func start(cfg *config.Config, logs *logger.Logs, db *sqlite.Database) {
	logs.Server.Info("Starting (2/2)...")

	router := chi.NewRouter()
	controlers.NewHand(logs, db).Register(router)
	logs.Server.Debug("Create router and use handlers")

	server := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      router,
		WriteTimeout: cfg.Server.Wto,
		ReadTimeout:  cfg.Server.Rto,
	}
	logs.Server.Debug("Create server")

	log.Print("Start")
	logs.Server.Info("Listen server")
	err := server.ListenAndServe()
	if err != nil {
		logs.Server.Error(fmt.Sprint("Stop server, error listen:", err))
	}
}

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
	cfg, err := config.Load() //Loading config - main/config/config.yaml. Use: cfg.<Name>
	if err != nil {
		log.Fatal("Error load config:", err)
	}

	logs, err := logger.New(cfg) //Loading logger. Use: logs.<Name>.<LevelMod>
	if err != nil {
		log.Fatal("Error load logger:", err)
	}

	db, err := sqlite.New(cfg, logs) //Loading sqlite database
	if err != nil {
		logs.Server.Warn("Stop server, error connect to database")
		logs.Sqlite.Error(fmt.Sprintf("Error connect database: %v", err))
		log.Fatal("Error")
	}

	logs.Sqlite.Info("Connect sqlite")
	logs.Server.Info("Starting (1/2)...")
	start(cfg, logs, db) //Start function to start the server
}

func start(cfg *config.Config, logs *logger.Logs, db *sqlite.Database) {
	logs.Server.Info("Starting (2/2)...")

	router := chi.NewRouter()                     //Create chi router
	controlers.NewHand(logs, db).Register(router) //Loading handlers
	logs.Server.Debug("Create router and use handlers")

	server := &http.Server{ //Create http server
		Addr:         cfg.Server.Addr,
		Handler:      router,
		WriteTimeout: cfg.Server.Wto,
		ReadTimeout:  cfg.Server.Rto,
	}
	logs.Server.Debug("Create server")

	log.Print("Start")
	logs.Server.Info("Listen server")
	err := server.ListenAndServe() //Listen
	if err != nil {
		logs.Server.Error(fmt.Sprint("Stop server, error listen:", err))
	}
}

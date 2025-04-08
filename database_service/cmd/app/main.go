package main

import (
	dbCfg "db_service/internal/configs/config_db"
	cfggo "db_service/internal/configs/config_go"
	"db_service/internal/storage/postgre"
	"db_service/models"
	"log/slog"
	"os"
)

func main() {
	// conf app
	cfg := cfggo.MustLoad()

	// log conf
	log := setupLogger(cfg.Env)
	log.Debug("logger start")

	// db initializing
	dbCfg, err := dbCfg.LoadConfigDB()
	if err != nil {
		log.Error("config databaz err")
		panic("database nil")
	}

	// db open
	db, err := postgre.NewStorage(dbCfg)
	if err != nil {
		log.Error("dont make database", db)
	}

	// db migrations
	if err = db.DB.AutoMigrate(&models.DatabaseModels{}); err != nil {
		log.Error("not get migrates ")
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

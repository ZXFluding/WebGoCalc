package main

import (
	"log/slog"
	"os"
	"test/services/config"
	"test/services/lib/sl"
	"test/services/repositories/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//init config: env
	cfg := config.MustLoad()

	//init logger: slog
	log := setupLogger(cfg.Env)

	//Logger is up
	log.Info("starting rep_cal", slog.String("env", cfg.Env))
	log.Debug("debug masseges are enabled")

	//TODO: init Repositories: postgrers
	repository, err := postgres.New(&cfg.Storage)
	if err != nil {
		log.Error("failed to init db: ", sl.Err(err))
		os.Exit(1)
	}
	_ = repository
	//TODO: init router: gin-gonic

	//TODO: init controllers

	//TODO: init Services

}

// setup Logger configuraton
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}
	return log
}

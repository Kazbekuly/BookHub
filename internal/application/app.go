package application

import (
	"BookHub"
	"BookHub/internal/config"
	"BookHub/internal/handler"
	"BookHub/internal/logging"
	"BookHub/internal/repository"
	"BookHub/internal/service"
	"fmt"
)

func Run() error {
	logging.Init()
	logger := logging.GetLogger()
	logger.Info("start application")
	cfg := config.GetConfig()
	srv := new(BookHub.Server)
	db, err := repository.NewConnectionDB(*cfg)
	if err != nil {
		logger.Fatal("error while connecting to database")
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	_ = fmt.Sprintf("Starting server...\nhttp://localhost%v/\n", ":"+cfg.Listen.Port)
	if err := srv.Run(cfg, handlers.InitRoutes()); err != nil {
		logger.Info("Error while starting application")
	}
	return err
}

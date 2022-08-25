package main

import (
	"github.com/Services/imanuz_service/api_service/api"
	"github.com/Services/imanuz_service/api_service/config"
	"github.com/Services/imanuz_service/api_service/pkg/logger"
	"github.com/Services/imanuz_service/api_service/services"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api-gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}

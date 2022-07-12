package main

import (
	"net"

	"github.com/Services/imanuz_service/data_service/config"
	pb "github.com/Services/imanuz_service/data_service/genproto/data_service"
	"github.com/Services/imanuz_service/data_service/pkg/db"
	"github.com/Services/imanuz_service/data_service/pkg/logger"
	"github.com/Services/imanuz_service/data_service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "Data")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
		logger.String("password", cfg.PostgresPassword))
	connDB, err := db.ConnectToDB(cfg)

	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	dataService := service.NewDataService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterDataServiceServer(s, dataService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}

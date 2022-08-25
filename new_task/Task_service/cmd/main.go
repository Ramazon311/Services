package main

import (
	"net"
	"github/Services/newpro/Task_service/config"
	pb "github/Services/newpro/Task_service/genproto/task_service"
	"github/Services/newpro/Task_service/pkg/db"
	"github/Services/newpro/Task_service/pkg/logger"
	"github/Services/newpro/Task_service/service"
	"github/Services/newpro/Task_service/service/grpc_client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "task-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	//pgStorage := storage.NewStoragePg(connDB)

	grpcClient, err := grpc_client.New(cfg)
	if err != nil {
		log.Error("error establishing grpc connection", logger.Error(err))
		return
	}

	taskService := service.NewTaskService(connDB, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()

	pb.RegisterTaskServiceServer(s, taskService)
	reflection.Register(s)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}

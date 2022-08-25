package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/Gorilla-services/user-service1/config"
	"github.com/Gorilla-services/user-service1/pkg/db"
	"github.com/Gorilla-services/user-service1/pkg/logger"
)
var repo *userRepo
func  TestMain(m *testing.M){
	cfg := config.Load()
	connDB ,err :=db.ConnectToDB(cfg)
	if err!=nil{
		log.Fatal("sqlx conect to db",logger.Error(err))
	}
	repo = NewUserRepo(connDB)
	os.Exit(m.Run())
}
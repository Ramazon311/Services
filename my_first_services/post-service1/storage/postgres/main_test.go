package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/Gorilla-services/post-service1/config"
	"github.com/Gorilla-services/post-service1/pkg/db"
	"github.com/Gorilla-services/post-service1/pkg/logger"
)
var repo *postRepo
func  TestMain(m *testing.M){
	cfg := config.Load()
	connDB ,err :=db.ConnectToDB(cfg)
	if err!=nil{
		log.Fatal("sqlx conect to db",logger.Error(err))
	}
	repo = NewPostRepo(connDB)
	os.Exit(m.Run())
}
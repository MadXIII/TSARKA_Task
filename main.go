package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/madxiii/tsarka_task/configs"
	"github.com/madxiii/tsarka_task/http"
	"github.com/madxiii/tsarka_task/repository"
	"github.com/madxiii/tsarka_task/repository/postgres"
	"github.com/madxiii/tsarka_task/repository/redis"
	"github.com/madxiii/tsarka_task/server"
	"github.com/madxiii/tsarka_task/service"
)

func main() {
	cfg, err := configs.Get()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	userDB, err := postgres.UserConn(cfg.Postgres)
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}
	countCache, hashCache := redis.NewClients(cfg.Redis)
	repos := repository.New(userDB, countCache, hashCache)
	services := service.New(*repos)
	api := http.NewAPI(*services)

	serv := new(server.Server)
	if err := serv.Run(viper.GetString("port"), api.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running server: %s", err.Error())
	}
}

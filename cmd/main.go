package main

import (
	backgo "github.com/ArtemKeety/back-go.git"
	"github.com/ArtemKeety/back-go.git/internal/database"
	"github.com/ArtemKeety/back-go.git/internal/handler"
	"github.com/ArtemKeety/back-go.git/internal/repository"
	"github.com/ArtemKeety/back-go.git/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := database.NewDB(database.ConfDb{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "root",
		Database: "authgo",
		SslMode:  "disable",
	})
	if err != nil {
		logrus.Fatal(err)
	}

	r := repository.NewRepository(db)
	ser := service.NewService(r)
	h := handler.NewHandler(ser)
	s := new(backgo.Server)

	if err := s.Run("localhost:8080", h.InitRouter()); err != nil {
		logrus.Fatal(err)
	}

}

package main

import (
	"fmt"
	backgo "github.com/ArtemKeety/back-go.git"
	"github.com/ArtemKeety/back-go.git/internal/handler"
	"github.com/ArtemKeety/back-go.git/internal/repository"
	"github.com/ArtemKeety/back-go.git/internal/service"
)

func main() {

	//todo db

	r := repository.NewRepository(nil)
	ser := service.NewService(r)
	h := handler.NewHandler(ser)
	s := new(backgo.Server)

	if err := s.Run("localhost:8080", h.InitRouter()); err != nil {
		fmt.Println(err)
	}

}

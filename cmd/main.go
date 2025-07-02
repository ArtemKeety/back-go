package main

import (
	"fmt"
	backgo "github.com/ArtemKeety/back-go.git"
	"github.com/ArtemKeety/back-go.git/internal/handler"
)

func main() {

	h := handler.NewHandler(nil)
	s := new(backgo.Server)

	if err := s.Run("localhost:8080", h.InitRouter()); err != nil {
		fmt.Println(err)
	}

}

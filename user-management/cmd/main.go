package main

import (
	"log"
	"user-management/pkg/handler"
	httpserver "user-management/pkg/http-server"
	"user-management/pkg/repository"
	"user-management/pkg/service"
)

func main() {

	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handler := handler.NewHandler(services)

	router := handler.InitRoutes()
	if err := httpserver.NewHttpServer(":8080", router); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"forum/web/delivery"
	"forum/web/server"
	"forum/web/service"
	"forum/web/storage"
)

func main() {
	db, err := storage.CreateDB()
	if err != nil {
		log.Fatal(err)
	}
	if err := storage.CreateTables(db); err != nil {
		log.Fatal(err)
	}
	storages := storage.NewStorage(db)
	services := service.NewService(storages)
	handlers := delivery.NewHandler(services)
	server := new(server.Server)
	if err := server.Start(handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

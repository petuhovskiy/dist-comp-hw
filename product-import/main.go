//go:generate swag init
package main

import (
	"log"
	"net/http"
	"product-import/config"
	"product-import/handlers"
	"product-import/routers"
	"product-import/service"
)

// @title Product Import API
// @version 1.0
// @description This is sample product import server, made as dist-comp homework.

// @contact.name Arthur Petukhovsky
// @contact.url https://t.me/petuhovskiy
// @contact.email petuhovskiy@yandex.ru

// @host localhost:8082
// @BasePath /
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf, err := config.EnvConfig()
	if err != nil {
		log.Panic("failed to read config, err=", err)
	}

	// clients
	sender := service.NewQueueSender(conf.AmqpURL, conf.QueueImport)

	// initializing handlers
	importHandler := handlers.NewImport(sender)
	handler := routers.CreateRouter(importHandler)

	log.Println("Serving at http://localhost" + conf.BindAddr)
	err = http.ListenAndServe(conf.BindAddr, handler)
	log.Println(err)
}

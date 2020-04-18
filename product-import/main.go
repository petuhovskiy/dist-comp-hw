//go:generate swag init
package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"lib/httputil"
	"lib/pb"
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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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

	// clients gRPC
	gconn, err := grpc.Dial(
		conf.AuthGrpc,
		grpc.WithInsecure(),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	authCli := pb.NewAuthClient(gconn)

	// clients
	sender := service.NewQueueSender(conf.AmqpURL, conf.QueueImport)

	// initializing handlers
	authMiddleware := httputil.AuthMiddleware(authCli)
	importHandler := handlers.NewImport(sender)
	handler := routers.CreateRouter(importHandler, authMiddleware)

	log.Println("Serving at http://localhost" + conf.BindAddr)
	err = http.ListenAndServe(conf.BindAddr, handler)
	log.Println(err)
}

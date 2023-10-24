package main

import (
	"context"
	"days-left/internal/handler"
	"days-left/internal/service"
	"days-left/pkg/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	//Load env
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error occured while loading env variables: ", err)
	}

	//Init services
	services := service.NewService()

	//Init handlers
	handler := handler.NewHandler(services)

	//Init server
	server := new(server.Server)

	//Run server
	go func() {
		if err := server.Run(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Days left service started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("Days left service shutting down")

	//Graceful shutdown
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Fatal("error occured while shutting down http server: ", err)
	}

}

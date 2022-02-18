package main

import (
	"os"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/GRPC/grpcClient"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/handler"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/database"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/repository"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/server"
	"stlab.itechart-group.com/go/food_delivery/authentication_service/service"
)

func main() {
	logger := logging.GetLogger()

	db, err := database.NewPostgresDB(database.PostgresDB{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		logger.Panicf("failed to initialize db:%s", err.Error())
	}

	grpcCli := grpcClient.NewGRPCClient()

	rep := repository.NewRepository(db, logger)
	ser := service.NewService(rep, grpcCli, logger)
	handlers := handler.NewHandler(logger, ser)

	port := os.Getenv("API_SERVER_PORT")
	serv := new(server.Server)

	if err := serv.Run(port, handlers.InitRoutes()); err != nil {
		logger.Panicf("Error occured while running http server: %s", err.Error())
	}

}

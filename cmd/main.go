package main

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/handler"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/database"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/repository"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/service"
	"os"
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

	rep := repository.NewRepository(db, logger)
	ser := service.NewService(rep, logger)
	handlers := handler.NewHandler(logger, ser)
	// Setup router
	router := handlers.InitRoutes()
	port := os.Getenv("API_SERVER_PORT")
	logger.Fatal(router.Run(":" + port))
}

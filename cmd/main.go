package main

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/handler"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/database"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/repository"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/service"
	"github.com/spf13/viper"
)

func main() {
	logger := logging.GetLogger()
	viper.AutomaticEnv()
	db, err := database.NewPostgresDB(database.PostgresDB{
		Host:     viper.GetString("HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_DATABASE"),
		SSLMode:  viper.GetString("DB_SSL_MODE"),
	})
	if err != nil {
		logger.Panicf("failed to initialize db:%s", err.Error())
	}

	rep := repository.NewRepository(db, logger)
	ser := service.NewService(rep, logger)
	handlers := handler.NewHandler(logger, ser)
	// Setup router
	router := handlers.InitRoutes()
	port := viper.GetString("API_SERVER_PORT")
	logger.Fatal(router.Run(":" + port))
}

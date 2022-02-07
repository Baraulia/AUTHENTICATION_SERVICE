package main

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/database"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/router"
	"github.com/spf13/viper"
	"log"
)

func main() {
	var err error

	// Setup database
	database.DB, err = database.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()

	// Setup router
	router := router.NewRoutes()
	port := viper.GetString("PORT")
	log.Fatal(router.Run(":" + port))
}


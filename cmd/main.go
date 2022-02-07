package main

import (
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/database"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/router"
	"github.com/spf13/viper"
	"log"
	"net/http"
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

<<<<<<< HEAD
=======
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello team! Auth service stage")
>>>>>>> 02fca46779994e8dc2241cbb0bb4dc18295bdf59
}


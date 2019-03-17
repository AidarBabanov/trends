package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
	"trends/constants"
	"trends/controllers"
	"trends/database"
	"trends/models"
)

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

func main() {
	// Loading environment variables
	log.Println("Loading environment variables.")
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	// Connecting to database
	_, err2 := database.Connect()
	if err2 != nil {
		log.Fatal(err2)
	}

	database.DB.AutoMigrate(models.Trends{})
	database.DB.AutoMigrate(models.GainingTrend{})

	// Setting application port
	port := os.Getenv("app_port")
	if port == "" {
		log.Fatal("$app_port not set")
	} else {
		log.Println(fmt.Sprintf("$app_port: %s", port))
	}

	// Setting router and controllers
	log.Println("Setting router and controllers.")
	router := mux.NewRouter()

	router.HandleFunc(constants.TrendsURI, controllers.Create).Methods(POST)
	router.HandleFunc(constants.TrendsURI, controllers.Get).Methods(GET)
	router.HandleFunc(constants.TrendsURI, controllers.Delete).Methods(DELETE)

	router.HandleFunc(constants.GainingURI, controllers.Create2).Methods(POST)
	router.HandleFunc(constants.GainingURI, controllers.Get2).Methods(GET)
	router.HandleFunc(constants.GainingURI, controllers.Delete2).Methods(DELETE)

	// Setting server
	log.Println("Router and controllers set successfully.")
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
	log.Println("Maintaining web application...")
}

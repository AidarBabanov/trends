package main

import (
	"./database"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
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

	//router.HandleFunc(c.RegisterURI, contr.Register).Methods(POST)

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

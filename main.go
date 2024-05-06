package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang_Project/models"
	"golang_Project/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create a new gin instance
	r := gin.Default()

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)

	// Load the routes
	routes.AuthRoutes(r)

	// Run the first server on port 8080
	go func() {
		fmt.Println("Starting server 1 on the port 8080....")
		r.Run(":8080")
	}()

	// Create a new mux router for the second server
	t := routes.Router()

	// Run the second server on port 8181
	fmt.Println("Starting server 2 on the port 8081....")
	log.Fatal(http.ListenAndServe(":8081", t))
}

/*
func main() {
	r := routes.Router()

	fmt.Println("Starting server on the port 8080....")

	log.Fatal(http.ListenAndServe(":8080", r))

} */

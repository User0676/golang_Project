package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"golang_Project/models"
	"log"
	"net/http"
	"os"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nill {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nill {
		panic(err)
	}

	err = db.Ping()

	if err != nill {
		panic(err)
	}

	fmt.Println("Succesfully connected to postgres")
	return db
}

func GetInstructor() {

}

func GetAllInstructor() {

}

func HireInstructor(w http.ResponseWriter, r *http.Request) {
	var instructor models.Instructor

	err := json.NewDecoder(r.Body).Decode(&instructor)

	if err != nill {
		log.Fatal("Unable to decode the request body. %v", err)
	}

	insertID := insertInstructor(instructor)

	res := response{
		ID:      insertID,
		Message: "instructor hired succesfully",
	}

	json.NewDecoder(w).Encode(res)

}

func UpdateInstructor() {

}

func FireInstructor() {

}

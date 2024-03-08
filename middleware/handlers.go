package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang_Project/models"
	"log"
	"net/http"
	"os"
	"strconv"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
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

func GetInstructor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nill {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	instructor, err := getInstructor(int64(id))

	if err != nill {
		log.Fatalf("unable to get instructor. %v", err)
	}

	json.NewDecoder(w).Encode(instructor)

}

func GetAllInstructor(w http.ResponseWriter, r *http.Request) {
	instructors, err := getAllInstructor()

	if err != nil {
		log.Fatalf("Unable to get all the shocks %v", err)
	}

	json.NewEncoder(w).Encode(instructors)
}

func HireInstructor(w http.ResponseWriter, r *http.Request) {
	var instructor models.Instructor

	err := json.NewDecoder(r.Body).Decode(&instructor)

	if err != nill {
		log.Fatal("Unable to decode the request body. %v", err)
	}

	insertID := hireInstructor(instructor)

	res := response{
		ID:      insertID,
		Message: "instructor hired succesfully",
	}

	json.NewDecoder(w).Encode(res)

}

func UpdateInstructor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	var instructor models.Instructor
	err = json.NewDecoder(r.Body).Decode(&instructor)

	if err != nil {
		log.Fatalf("Unable to decode request body. %v", err)
	}

	updatedRows := updateInstructor(int64(id), instructor)

	msg := fmt.Sprintf("instructor updated succesfully. Total rows/records affected %v", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewDecoder(w).Encode(res)
}

func FireInstructor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}
	deletedRows := fireInstructor(int64(int))
	msg := fmt.Sprintf("Instructor fired successfully. Total rows/records %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func hireInstructor(instructor models.Instructor) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO Instructors(instructor_name,specialization,gender) VALUES ($1, $2, $3) RETURNING instructor_id`
	var id int64

	err := db.QueryRow(sqlStatement, instructor.InstructorName, instructor.Specialization, instructor.Gender).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record. %v", err)
	return id
}
func getInstructor(id int64) (models.Instructor, error) {
	db := createConnection()

	defer db.Close()

	var instructor models.Instructor

	sqlStatement := `SELECT * FROM Instructors WHERE instructor_id=$1`

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&instructor.InstructorId, &instructor.InstructorName, &instructor.Specialization, &instructor.Gender)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned!")
		return instructor, nil
	case nil:
		return instructor, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return instructor, err

}
func getAllInstructor() ([]models.Instructors, error) {
	db := createConnection()
	defer db.Close()
	var instructors []models.Instructor

	sqlStatement := `SELECT * FROM Instructors`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var instructor models.Instructor

		err = rows.Scan(&instructor.InstructorId, &instructor.InstructorName, &instructor.Specialization, &instructor.Gender)

		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		instructors = append(instructors, instructor)
	}
	return instructors, err

}

func fireInstructor(id int64) int64 {

}
func updateInstructor(id int64, instructor models.Instructor) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE Instructors SET instructor_name=$2,specialization=$3,gender=$4 WHERE instructor_id = $1 `

	db.Exec(sqlStatement, id, instructor.InstructorName, instructor.Specialization, instructor.Gender)
}

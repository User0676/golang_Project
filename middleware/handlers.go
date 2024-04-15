package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully connected to postgres")
	return db
}

func GetInstructor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	instructor, err := getInstructor(int64(id))

	if err != nil {
		log.Fatalf("unable to get instructor. %v", err)
	}

	json.NewEncoder(w).Encode(instructor)

}

func GetAllInstructor(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for pagination
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")
	sortBy := r.URL.Query().Get("sortBy")
	sortOrder := r.URL.Query().Get("sortOrder")
	filter := r.URL.Query().Get("filter")

	// Default values
	page := 1
	pageSize := 10

	// Convert page and pageSize parameters to integers
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	instructors, err := getAllInstructorPaginated(page, pageSize, sortBy, sortOrder, filter)

	if err != nil {
		log.Fatalf("Unable to get paginated instructors: %v", err)
	}

	json.NewEncoder(w).Encode(instructors)
}

func getAllInstructorPaginated(page int, pageSize int, sortBy string, sortOrder string, filter string) ([]models.Instructor, error) {
	db := createConnection()
	defer db.Close()
	var instructors []models.Instructor

	// Calculate offset
	offset := (page - 1) * pageSize

	// Build SQL query with filtering and sorting
	sqlStatement := fmt.Sprintf("SELECT * FROM Instructors")

	// Add WHERE clause for filtering
	if filter != "" {
		sqlStatement += fmt.Sprintf(" WHERE instructor_name LIKE '%%%s%%' OR specialization LIKE '%%%s%%'", filter, filter)
	}

	// Add ORDER BY clause for sorting
	if sortBy != "" {
		sqlStatement += fmt.Sprintf(" ORDER BY %s %s", sortBy, sortOrder)
	}

	// Add LIMIT and OFFSET clauses for pagination
	sqlStatement += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var instructor models.Instructor

		err = rows.Scan(&instructor.InstructorId, &instructor.InstructorName, &instructor.Specialization, &instructor.Gender)

		if err != nil {
			log.Fatalf("Unable to scan the row: %v", err)
		}
		instructors = append(instructors, instructor)
	}
	return instructors, err
}

func HireInstructor(w http.ResponseWriter, r *http.Request) {
	var instructor models.Instructor

	err := json.NewDecoder(r.Body).Decode(&instructor)

	if err != nil {
		log.Fatal("Unable to decode the request body. %v", err)
	}

	insertID := hireInstructor(instructor)

	res := response{
		ID:      insertID,
		Message: "instructor hired succesfully",
	}

	json.NewEncoder(w).Encode(res)

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

	json.NewEncoder(w).Encode(res)
}

func FireInstructor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}
	deletedRows := fireInstructor(int64(id))
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
func getAllInstructor() ([]models.Instructor, error) {
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

func updateInstructor(id int64, instructor models.Instructor) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE Instructors SET instructor_name=$2,specialization=$3,gender=$4 WHERE instructor_id = $1 `

	res, err := db.Exec(sqlStatement, id, instructor.InstructorName, instructor.Specialization, instructor.Gender)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the rows %v", err)
	}

	fmt.Printf("Total rows affected %v", rowsAffected)
	return rowsAffected
}

func fireInstructor(id int64) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM Instructors WHERE instructor_id = $1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the rows %v", err)
	}

	fmt.Printf("Total rows affected %v", rowsAffected)
	return rowsAffected
}

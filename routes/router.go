package routes

import (
	"github.com/gorilla/mux"
	"golang_Project/middlewares"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/instructors/{id}", middlewares.GetInstructor).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/instructors", middlewares.GetAllInstructor).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/hireinstructors", middlewares.HireInstructor).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/updateinstructors/{id}", middlewares.UpdateInstructor).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/fireinstructors/{id}", middlewares.FireInstructor).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/instructors", middlewares.GetAllInstructor).Methods("GET", "OPTIONS").Queries("page", "{page:[0-9]+}", "pageSize", "{pageSize:[0-9]+}")
	router.HandleFunc("/api/instructors", middlewares.GetAllInstructor).Methods("GET", "OPTIONS").Queries("page", "{page:[0-9]+}", "pageSize", "{pageSize:[0-9]+}", "sortBy", "{sortBy}", "sortOrder", "{sortOrder}", "filter", "{filter}")

	return router
}

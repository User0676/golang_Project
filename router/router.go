package router

import (
	"github.com/gorilla/mux"
	"golang_Project/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/instructors/{id}", middleware.GetInstructor).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/instructors", middleware.GetAllInstructor).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/hireinstructors", middleware.HireInstructor).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/updateinstructors/{id}", middleware.UpdateInstructor).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/fireinstructors/{id}", middleware.FireInstructor).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/instructors", middleware.GetAllInstructor).Methods("GET", "OPTIONS").Queries("page", "{page:[0-9]+}", "pageSize", "{pageSize:[0-9]+}", "sortBy", "{sortBy}", "sortOrder", "{sortOrder}", "filter", "{filter}")

	return router
}

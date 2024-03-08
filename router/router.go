package router

import (
	"github.com/gorilla/mux"
	"golang_Project/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/instructor/{id}", middleware.GetInstructor).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/instructor", middleware.GetAllInstructor).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/hireinstructor", middleware.HireInstructor).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/update/{id}", middleware.UpdateInstructor).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/fireinstructor/{id}", middleware.FireInstructor).Methods("DELETE", "OPTIONS")

	return router
}

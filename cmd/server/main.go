package main

import (
	"log"
	"net/http"

	"github.com/0xivanov/backend-api/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	h := handlers.NewHandlers()

	// Car routes
	router.HandleFunc("/cars", h.GetAllCars).Methods("GET")
	router.HandleFunc("/cars", h.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", h.GetCarById).Methods("GET")
	router.HandleFunc("/cars/{id}", h.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", h.DeleteCarById).Methods("DELETE")

	// Garage routes
	router.HandleFunc("/garages/dailyAvailabilityReport", h.GetGarageReport).Methods("GET")
	router.HandleFunc("/garages", h.GetAllGarages).Methods("GET")
	router.HandleFunc("/garages", h.CreateGarage).Methods("POST")
	router.HandleFunc("/garages/{id}", h.GetGarageById).Methods("GET")
	router.HandleFunc("/garages/{id}", h.UpdateGarage).Methods("PUT")
	router.HandleFunc("/garages/{id}", h.DeleteGarageById).Methods("DELETE")

	// Maintenance routes
	router.HandleFunc("/maintenance/monthlyRequestsReport", h.MonthlyRequestsReport).Methods("GET")
	router.HandleFunc("/maintenance", h.GetAllMaintenances).Methods("GET")
	router.HandleFunc("/maintenance", h.CreateMaintenance).Methods("POST")
	router.HandleFunc("/maintenance/{id}", h.GetMaintenanceById).Methods("GET")
	router.HandleFunc("/maintenance/{id}", h.UpdateMaintenance).Methods("PUT")
	router.HandleFunc("/maintenance/{id}", h.DeleteMaintenance).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Printf("Server starting on port 8088...")
	if err := http.ListenAndServe(":8088", handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

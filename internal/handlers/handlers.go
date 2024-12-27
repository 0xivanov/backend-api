package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/0xivanov/backend-api/internal/models"

	"github.com/gorilla/mux"
)

type Handlers struct {
	cars         map[int64]models.Car
	garages      map[int64]models.Garage
	maintenances map[int64]models.Maintenance
	nextID       map[string]int64
}

func NewHandlers() *Handlers {
	return &Handlers{
		cars:         make(map[int64]models.Car),
		garages:      make(map[int64]models.Garage),
		maintenances: make(map[int64]models.Maintenance),
		nextID: map[string]int64{
			"car":         1,
			"garage":      1,
			"maintenance": 1,
		},
	}
}

// Car handlers
func (h *Handlers) GetAllCars(w http.ResponseWriter, r *http.Request) {
	carMake := r.URL.Query().Get("carMake")
	garageIDStr := r.URL.Query().Get("garageId")
	fromYearStr := r.URL.Query().Get("fromYear")
	toYearStr := r.URL.Query().Get("toYear")

	cars := make([]models.Car, 0)
	for _, car := range h.cars {
		if carMake != "" && car.Make != carMake {
			continue
		}

		if garageIDStr != "" {
			garageID, _ := strconv.ParseInt(garageIDStr, 10, 64)
			found := false
			for _, id := range car.GarageIDs {
				if id == garageID {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		if fromYearStr != "" {
			fromYear, _ := strconv.ParseInt(fromYearStr, 10, 32)
			if int32(fromYear) > car.ProductionYear {
				continue
			}
		}

		if toYearStr != "" {
			toYear, _ := strconv.ParseInt(toYearStr, 10, 32)
			if int32(toYear) < car.ProductionYear {
				continue
			}
		}

		cars = append(cars, car)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func (h *Handlers) CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(car.GarageIDs) == 0 {
		http.Error(w, "garage ids are missing", http.StatusBadRequest)
		return
	}
	car.Garages = make([]models.Garage, 0)
	for _, garageId := range car.GarageIDs {
		if _, ok := h.garages[garageId]; !ok {
			http.Error(w, "garage does not exist", http.StatusBadRequest)
			return
		}
		car.Garages = append(car.Garages, h.garages[garageId])
	}

	car.ID = h.nextID["car"]
	h.nextID["car"]++
	h.cars[car.ID] = car

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func (h *Handlers) GetCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	car, exists := h.cars[id]
	if !exists {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func (h *Handlers) UpdateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := h.cars[id]; !exists {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car.ID = id
	h.cars[id] = car

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func (h *Handlers) DeleteCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := h.cars[id]; !exists {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	delete(h.cars, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

// Garage handlers
func (h *Handlers) GetAllGarages(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	garages := make([]models.Garage, 0)
	for _, garage := range h.garages {
		if city != "" && garage.City != city {
			continue
		}
		garages = append(garages, garage)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(garages)
}

func (h *Handlers) CreateGarage(w http.ResponseWriter, r *http.Request) {
	var garage models.Garage
	if err := json.NewDecoder(r.Body).Decode(&garage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	garage.ID = h.nextID["garage"]
	h.nextID["garage"]++
	h.garages[garage.ID] = garage

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(garage)
}

func (h *Handlers) GetGarageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	garage, exists := h.garages[id]
	if !exists {
		http.Error(w, "Garage not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(garage)
}

func (h *Handlers) UpdateGarage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := h.garages[id]; !exists {
		http.Error(w, "Garage not found", http.StatusNotFound)
		return
	}

	var garage models.Garage
	if err := json.NewDecoder(r.Body).Decode(&garage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	garage.ID = id
	h.garages[id] = garage

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(garage)
}

func (h *Handlers) DeleteGarageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := h.garages[id]; !exists {
		http.Error(w, "Garage not found", http.StatusNotFound)
		return
	}

	delete(h.garages, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

// Maintenance handlers
func (h *Handlers) GetAllMaintenances(w http.ResponseWriter, r *http.Request) {
	carIDStr := r.URL.Query().Get("carId")
	garageIDStr := r.URL.Query().Get("garageId")
	startDateStr := r.URL.Query().Get("startDate")
	endDateStr := r.URL.Query().Get("endDate")

	maintenances := make([]models.Maintenance, 0)
	for _, maintenance := range h.maintenances {
		if carIDStr != "" {
			if maintenance.CarID != carIDStr {
				continue
			}
		}

		if garageIDStr != "" {
			if maintenance.GarageID != garageIDStr {
				continue
			}
		}

		if startDateStr != "" {
			startDate, _ := time.Parse("2006-01-02", startDateStr)
			if maintenance.ScheduledDate.Before(startDate) {
				continue
			}
		}

		if endDateStr != "" {
			endDate, _ := time.Parse("2006-01-02", endDateStr)
			if maintenance.ScheduledDate.After(endDate) {
				continue
			}
		}

		maintenances = append(maintenances, maintenance)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maintenances)
}

func (h *Handlers) CreateMaintenance(w http.ResponseWriter, r *http.Request) {
	var maintenance models.Maintenance
	if err := json.NewDecoder(r.Body).Decode(&maintenance); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	maintenance.ID = h.nextID["maintenance"]
	h.nextID["maintenance"]++
	h.maintenances[maintenance.ID] = maintenance

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maintenance)
}

func (h *Handlers) GetMaintenanceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	maintenance, exists := h.maintenances[id]
	if !exists {
		http.Error(w, "Maintenance not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maintenance)
}

func (h *Handlers) UpdateMaintenance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := h.maintenances[id]; !exists {
		http.Error(w, "Maintenance not found", http.StatusNotFound)
		return
	}

	var maintenance models.Maintenance
	if err := json.NewDecoder(r.Body).Decode(&maintenance); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// The date is already parsed by UnmarshalJSON in the models.Maintenance
	maintenance.ID = id
	h.maintenances[id] = maintenance

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maintenance)
}

func (h *Handlers) DeleteMaintenance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := h.maintenances[id]; !exists {
		http.Error(w, "Maintenance not found", http.StatusNotFound)
		return
	}

	delete(h.maintenances, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

// Report handlers
func (h *Handlers) GetGarageReport(w http.ResponseWriter, r *http.Request) {
	garageIDStr := r.URL.Query().Get("garageId")
	startDateStr := r.URL.Query().Get("startDate")
	endDateStr := r.URL.Query().Get("endDate")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		http.Error(w, "Invalid start date", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		http.Error(w, "Invalid end date", http.StatusBadRequest)
		return
	}

	garageID, err := strconv.ParseInt(garageIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid garage ID", http.StatusBadRequest)
		return
	}

	garage, exists := h.garages[garageID]
	if !exists {
		http.Error(w, "Garage not found", http.StatusNotFound)
		return
	}

	report := make([]models.GarageDailyAvailabilityReport, 0)
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		requests := 0
		for _, m := range h.maintenances {
			if m.GarageID == garageIDStr && m.ScheduledDate.Equal(d) {
				requests++
			}
		}

		report = append(report, models.GarageDailyAvailabilityReport{
			Date:              models.CustomTime{Time: d},
			Requests:          int32(requests),
			AvailableCapacity: garage.Capacity - int32(requests),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func (h *Handlers) MonthlyRequestsReport(w http.ResponseWriter, r *http.Request) {
	garageIDStr := r.URL.Query().Get("garageId")
	startMonth := r.URL.Query().Get("startMonth")
	endMonth := r.URL.Query().Get("endMonth")

	startDate, err := time.Parse("2006-01", startMonth)
	if err != nil {
		http.Error(w, "Invalid start month", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01", endMonth)
	if err != nil {
		http.Error(w, "Invalid end month", http.StatusBadRequest)
		return
	}

	var report []models.MonthlyRequestsReport
	currentDate := startDate
	for !currentDate.After(endDate) {
		requests := 0
		for _, m := range h.maintenances {
			if m.GarageID == garageIDStr &&
				m.ScheduledDate.Year() == currentDate.Year() &&
				m.ScheduledDate.Month() == currentDate.Month() {
				requests++
			}
		}

		monthReport := models.MonthlyRequestsReport{}
		monthReport.YearMonth.Year = int32(currentDate.Year())
		monthReport.YearMonth.Month = currentDate.Month().String()
		monthReport.YearMonth.MonthValue = int32(currentDate.Month())
		monthReport.YearMonth.LeapYear = currentDate.Year()%4 == 0
		monthReport.Requests = int32(requests)

		report = append(report, monthReport)
		currentDate = currentDate.AddDate(0, 1, 0)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

package models

import (
	"time"
)

type Car struct {
	ID             int64    `json:"id"`
	Make           string   `json:"make"`
	Model          string   `json:"model"`
	ProductionYear int32    `json:"productionYear"`
	LicensePlate   string   `json:"licensePlate"`
	GarageIDs      []int64  `json:"garageIds,omitempty"`
	Garages        []Garage `json:"garages,omitempty"`
}

type Garage struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	City     string `json:"city"`
	Capacity int32  `json:"capacity"`
}

type Maintenance struct {
	ID            int64      `json:"id"`
	CarID         string     `json:"carId"`
	CarName       string     `json:"carName,omitempty"`
	ServiceType   string     `json:"serviceType"`
	ScheduledDate CustomTime `json:"scheduledDate"`
	GarageID      string     `json:"garageId"`
	GarageName    string     `json:"garageName,omitempty"`
}

type CustomTime struct {
	time.Time
}

type MonthlyRequestsReport struct {
	YearMonth struct {
		Year       int32  `json:"year"`
		Month      string `json:"month"`
		LeapYear   bool   `json:"leapYear"`
		MonthValue int32  `json:"monthValue"`
	} `json:"yearMonth"`
	Requests int32 `json:"requests"`
}

type GarageDailyAvailabilityReport struct {
	Date              CustomTime `json:"date" format:"2006-01-02"`
	Requests          int32      `json:"requests"`
	AvailableCapacity int32      `json:"availableCapacity"`
}

const timeFormat = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	t, err := time.Parse(timeFormat, str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

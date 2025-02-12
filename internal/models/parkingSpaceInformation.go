package models

import (
	"time"
)

type ParkingSpaceInformation struct {
	ID              int           `json:"id" db:"id"`
	VehicleID       int           `json:"vehicle_id" db:"vehicle_id"`
	ParkingSpaceID  int           `json:"parking_space_id" db:"parking_space_id"`
	ArrivalTime     time.Time     `json:"arrival_time" db:"arrival_time"`
	DepartureTime   time.Time     `json:"departure_time" db:"departure_time"`
	ParkingDuration time.Duration `json:"parking_duration" db:"parking_duration"`
}

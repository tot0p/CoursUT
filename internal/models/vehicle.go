package models

type Vehicle struct {
	ID          int         `json:"id" db:"id"`
	Plate       string      `json:"plate" db:"plate"`
	VehicleType VehicleType `json:"vehicleType" db:"vehicle_type"`
}

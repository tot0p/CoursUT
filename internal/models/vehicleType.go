package models

import (
	"time"
)

type VehicleType int

const (
	Unknown VehicleType = iota
	Car
	Truck
)

func IsValidVehicleType(value int) bool {
	return value >= int(Unknown) && value <= int(Truck)
}

func GetVehicleHourPrice(vehicleType VehicleType) int {
	switch vehicleType {
	case Car:
		return 5
	case Truck:
		return 12
	default:
		return 10
	}
}

func GetDegresiveParkingPrice(vehicleType VehicleType, duration time.Duration) int {
	// the price is degresive the more hour you stay
	price := GetVehicleHourPrice(vehicleType)
	// degresive coefficient
	coef := 1.0
	hours := int(duration.Hours())
	coef = 1.0 - (float64(hours) / 30.0)
	if coef < 0.5 {
		coef = 0.5
	}
	return int(float64(price*hours) * coef)

}

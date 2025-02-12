package models

type VehicleType int

const (
	Unknown VehicleType = iota
	Car
	Truck
)

func IsValidVehicleType(value int) bool {
	return value >= int(Unknown) && value <= int(Truck)
}

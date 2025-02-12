package api

import "github.com/tot0p/CoursUT/internal/models"

// VehicleInput is the input for the AddVehicleHandler
type VehicleInput struct {
	Plate string             `json:"plate"`
	Type  models.VehicleType `json:"vehicleType"`
}

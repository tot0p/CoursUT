package api

// VehicleInput is the input for the AddVehicleHandler
type VehicleInput struct {
	Plate string `json:"plate"`
	Type  string `json:"type"`
}

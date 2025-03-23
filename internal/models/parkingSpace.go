package models

import (
	"regexp"
)

type ParkingSpace struct {
	ID          int         `json:"id" db:"id"`
	VehicleType VehicleType `json:"vehicle_type" db:"vehicle_type"`
	SpaceNumber string      `json:"space_number" db:"space_number"`
}

// Format parking letter and 1 - 3 digits number (ex: A001, B123, C999)
const regexParkingSpace = `^[A-Z]{1}[0-9]{1,3}$`

// IsValidParkingSpace checks if the parking space number is valid
// A valid parking space number is a letter followed by 1 to 3 digits
// (ex: A001, B123, C999)
func IsValidParkingSpace(spaceNumber string) bool {
	return regexp.MustCompile(regexParkingSpace).MatchString(spaceNumber)
}

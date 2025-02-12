package models

type ParkingSpace struct {
	ID          int    `json:"id" db:"id"`
	SpaceNumber string `json:"space_number" db:"space_number"`
}

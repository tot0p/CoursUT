package api

// ReservationInput is the input for the AddReservationHandler
type ReservationInput struct {
	VehicleID           int    `json:"vehicle_id"`
	ReservationDuration string `json:"reservation_time"`
}

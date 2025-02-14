package parkingSpace

import (
	"context"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/models"
	"testing"
)

func TestCreateParkingSpace(t *testing.T) {
	err := database.InitDatabase()
	if err != nil {
		t.Error(err)
	}
	var ParkingSpace = models.ParkingSpace{
		SpaceNumber: "A32",
	}
	_, err = CreateParkingSpace(&ParkingSpace)
	if err != nil {
		t.Error(err)
	}
	// check if the parking space was created
	if ParkingSpace.SpaceNumber == "" {
		t.Error("Parking space was not created")
	}
	// query the parking space from the database
	p := models.ParkingSpace{}
	err = database.Conn.QueryRowContext(context.Background(), "SELECT * FROM parking_space WHERE id = ?;", ParkingSpace.ID).Scan(&p.ID, &p.SpaceNumber)
	if err != nil {
		t.Error(err)
	}
	// check if the parking space is the same
	if p != ParkingSpace {
		t.Error("Parking space is not the same")
	}
}

package parkingSpace

import (
	"context"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/models"
	"testing"
)

func TestCreateParkingSpace(t *testing.T) {
	err := database.InitDatabase()
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			t.Error(err)
		}
	}()
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

func TestCreateParkingSpaceAlreadyExists(t *testing.T) {
	err := database.InitDatabase()
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			t.Error(err)
		}
	}()
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
	_, err = CreateParkingSpace(&ParkingSpace)
	if err == nil {
		t.Error("Parking space was created twice")
	}
}

func TestGetParkingSpace(t *testing.T) {
	err := database.InitDatabase()
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			t.Error(err)
		}
	}()
	if err != nil {
		t.Error(err)
	}
	var ParkingSpace = models.ParkingSpace{
		SpaceNumber: "A32",
	}
	temp, err := CreateParkingSpace(&ParkingSpace)
	if err != nil {
		t.Error(err)
	}
	ParkingSpace = *temp
	p, err := GetParkingSpace(ParkingSpace.ID)
	if err != nil {
		t.Error(err)
	}
	if p.SpaceNumber != ParkingSpace.SpaceNumber {
		t.Error("Parking space is not the same")
	}
	if p.ID != ParkingSpace.ID {
		t.Error("Parking space is not the same")
	}
}

func TestGetParkingSpaceNotFound(t *testing.T) {
	err := database.InitDatabase()
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			t.Error(err)
		}
	}()
	if err != nil {
		t.Error(err)
	}
	_, err = GetParkingSpace(1)
	if err == nil {
		t.Error("Parking space was found")
	}
}

func TestGetParkingSpaces(t *testing.T) {
	err := database.InitDatabase()
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			t.Error(err)
		}
	}()
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
	parkingSpaces, err := GetParkingSpaces()
	if err != nil {
		t.Error(err)
	}
	if len(parkingSpaces) == 0 {
		t.Error("No parking spaces found")
	}
}

func TestDeleteParkingSpace(t *testing.T) {
	err := database.InitDatabase()
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			t.Error(err)
		}
	}()
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
	err = DeleteParkingSpace(ParkingSpace.ID)
	if err != nil {
		t.Error(err)
	}
	_, err = GetParkingSpace(ParkingSpace.ID)
	if err == nil {
		t.Error("Parking space was not deleted")
	}
}

func TestDeleteParkingSpaceNotFound(t *testing.T) {
	err := database.InitDatabase()
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			t.Error(err)
		}
	}()
	if err != nil {
		t.Error(err)
	}
	err = DeleteParkingSpace(1)
	if err == nil {
		t.Error("no error returned")
	}
}

func TestUpdateParkingSpace(t *testing.T) {
	err := database.InitDatabase()
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			t.Error(err)
		}
	}()
	if err != nil {
		t.Error(err)
	}
	var ParkingSpace = models.ParkingSpace{
		SpaceNumber: "A32",
	}
	temp, err := CreateParkingSpace(&ParkingSpace)
	if err != nil {
		t.Error(err)
	}
	ParkingSpace.ID = temp.ID
	ParkingSpace.SpaceNumber = "A33"
	err = UpdateParkingSpace(&ParkingSpace)
	if err != nil {
		t.Error(err)
	}
	p, err := GetParkingSpace(ParkingSpace.ID)
	if err != nil {
		t.Error(err)
	}
	if p.ID != ParkingSpace.ID {
		t.Error("Parking space was not updated")
	}
	if p.SpaceNumber != "A33" {
		t.Error("Parking space was not updated")
	}
}

func TestUpdateParkingSpaceNotFound(t *testing.T) {
	err := database.InitDatabase()
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			t.Error(err)
		}
	}()
	if err != nil {
		t.Error(err)
	}
	var ParkingSpace = models.ParkingSpace{
		ID:          1,
		SpaceNumber: "A32",
	}
	err = UpdateParkingSpace(&ParkingSpace)
	if err == nil {
		t.Error("no error returned")
	}
}

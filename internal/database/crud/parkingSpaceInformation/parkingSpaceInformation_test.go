package parkingSpaceInformation

import (
	"context"
	"fmt"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpace"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
	"github.com/tot0p/CoursUT/internal/models"
	"testing"
	"time"
)

func TestCreateParkingSpaceInformation(t *testing.T) {
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
	retParkingSpace, err := parkingSpace.CreateParkingSpace(&ParkingSpace)
	if err != nil {
		t.Error(err)
	}
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	retVehicle, err := vehicle.CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	var ParkingSpaceInformation = models.ParkingSpaceInformation{
		ParkingSpaceID:  retParkingSpace.ID,
		VehicleID:       retVehicle.ID,
		ArrivalTime:     time.Now(),
		DepartureTime:   time.Now().Add(time.Hour),
		ParkingDuration: time.Hour,
	}
	retParkingSpaceInformation, err := CreateParkingSpaceInformation(&ParkingSpaceInformation)
	if err != nil {
		t.Error(err)
	}
	// check if the parking space information was created
	if retParkingSpaceInformation.ID == 0 {
		t.Error("Parking space information was not created")
	}
	// query the parking space information from the database
	p := models.ParkingSpaceInformation{}
	err = database.Conn.QueryRowContext(context.Background(), "SELECT * FROM parking_space_information WHERE id = ?;", retParkingSpaceInformation.ID).Scan(&p.ID, &p.ParkingSpaceID, &p.VehicleID, &p.ArrivalTime, &p.DepartureTime, &p.ParkingDuration)
	if err != nil {
		t.Error(err)
	}
	// check if the parking space information is the same
	if p.ID != retParkingSpaceInformation.ID {
		t.Error("Parking space information is not the same")
	}
}

func TestGetParkingSpaceInformation(t *testing.T) {
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
	retParkingSpace, err := parkingSpace.CreateParkingSpace(&ParkingSpace)
	if err != nil {
		t.Error(err)
	}
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	retVehicle, err := vehicle.CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	var ParkingSpaceInformation = models.ParkingSpaceInformation{
		ParkingSpaceID:  retParkingSpace.ID,
		VehicleID:       retVehicle.ID,
		ArrivalTime:     time.Now(),
		DepartureTime:   time.Now().Add(time.Hour),
		ParkingDuration: time.Hour,
	}
	retParkingSpaceInformation, err := CreateParkingSpaceInformation(&ParkingSpaceInformation)
	if err != nil {
		t.Error(err)
	}
	retParkingSpaceInformation2, err := GetParkingSpaceInformation(retParkingSpaceInformation.ID)
	if err != nil {
		t.Error(err)
	}
	if retParkingSpaceInformation.ID != retParkingSpaceInformation2.ID {
		t.Error("Parking space information is not the same")
	}
}

func TestGetParkingSpaceInformationNotFound(t *testing.T) {
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
	_, err = GetParkingSpaceInformation(0)
	if err == nil {
		t.Error("Parking space information was found")
	}
}

func TestGetParkingSpaceInformations(t *testing.T) {
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
	retParkingSpace, err := parkingSpace.CreateParkingSpace(&ParkingSpace)
	if err != nil {
		t.Error(err)
	}
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	retVehicle, err := vehicle.CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	var ParkingSpaceInformation = models.ParkingSpaceInformation{
		ParkingSpaceID:  retParkingSpace.ID,
		VehicleID:       retVehicle.ID,
		ArrivalTime:     time.Now(),
		DepartureTime:   time.Now().Add(time.Hour),
		ParkingDuration: time.Hour,
	}
	_, err = CreateParkingSpaceInformation(&ParkingSpaceInformation)
	if err != nil {
		t.Error(err)
	}
	retParkingSpaceInformations, err := GetParkingSpaceInformations()
	if err != nil {
		t.Error(err)
	}
	if len(retParkingSpaceInformations) == 0 {
		t.Error("No parking space information found")
	}
}

func TestDeleteParkingSpaceInformation(t *testing.T) {
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
	retParkingSpace, err := parkingSpace.CreateParkingSpace(&ParkingSpace)
	if err != nil {
		t.Error(err)
	}
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	retVehicle, err := vehicle.CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	var ParkingSpaceInformation = models.ParkingSpaceInformation{
		ParkingSpaceID:  retParkingSpace.ID,
		VehicleID:       retVehicle.ID,
		ArrivalTime:     time.Now(),
		DepartureTime:   time.Now().Add(time.Hour),
		ParkingDuration: time.Hour,
	}
	retParkingSpaceInformation, err := CreateParkingSpaceInformation(&ParkingSpaceInformation)
	if err != nil {
		t.Error(err)
	}
	err = DeleteParkingSpaceInformation(retParkingSpaceInformation.ID)
	if err != nil {
		t.Error(err)
	}
	_, err = GetParkingSpaceInformation(retParkingSpaceInformation.ID)
	if err == nil {
		t.Error("Parking space information was found")
	}
}

func TestDeleteParkingSpaceInformationNotFound(t *testing.T) {
	err := DeleteParkingSpaceInformation(0)
	if err == nil {
		t.Error("Parking space information was found")
	}
}

func TestUpdateParkingSpaceInformation(t *testing.T) {
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
	retParkingSpace, err := parkingSpace.CreateParkingSpace(&ParkingSpace)
	if err != nil {
		t.Error(err)
	}
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	retVehicle, err := vehicle.CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	var ParkingSpaceInformation = models.ParkingSpaceInformation{
		ParkingSpaceID:  retParkingSpace.ID,
		VehicleID:       retVehicle.ID,
		ArrivalTime:     time.Now(),
		DepartureTime:   time.Now().Add(time.Hour),
		ParkingDuration: time.Hour,
	}
	retParkingSpaceInformation, err := CreateParkingSpaceInformation(&ParkingSpaceInformation)
	if err != nil {
		t.Error(err)
	}
	retParkingSpaceInformation.ArrivalTime = time.Now().Add(time.Hour)
	retParkingSpaceInformation.DepartureTime = time.Now().Add(2 * time.Hour)
	retParkingSpaceInformation.ParkingDuration = 2 * time.Hour
	err = UpdateParkingSpaceInformation(retParkingSpaceInformation)
	if err != nil {
		t.Error(err)
	}
	retParkingSpaceInformation2, err := GetParkingSpaceInformation(retParkingSpaceInformation.ID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(retParkingSpaceInformation.ArrivalTime.Unix())
	fmt.Println(retParkingSpaceInformation2.ArrivalTime.Unix())
	fmt.Println(retParkingSpaceInformation.ID)
	fmt.Println(retParkingSpaceInformation2.ID)
	if retParkingSpaceInformation2.ID != retParkingSpaceInformation.ID {
		t.Error("Parking space information is not the same")
	}
	if retParkingSpaceInformation2.ArrivalTime.Unix() != retParkingSpaceInformation.ArrivalTime.Unix() {
		t.Error("Parking space information is not the same")
	}
	if retParkingSpaceInformation2.DepartureTime.Unix() != retParkingSpaceInformation.DepartureTime.Unix() {
		t.Error("Parking space information is not the same")
	}
	if retParkingSpaceInformation2.ParkingDuration != retParkingSpaceInformation.ParkingDuration {
		t.Error("Parking space information is not the same")
	}
}

func TestUpdateParkingSpaceInformationNotFound(t *testing.T) {
	err := database.InitDatabase()
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			t.Error(err)
		}
	}()
	err = UpdateParkingSpaceInformation(&models.ParkingSpaceInformation{})
	if err == nil {
		t.Error("Parking space information was found")
	}
}

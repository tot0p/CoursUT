package vehicle

import (
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/models"
	"testing"
)

func TestCreateVehicle(t *testing.T) {
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
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	retVehicle, err := CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	retVehicle2, err := GetVehicle(retVehicle.ID)
	if err != nil {
		t.Error(err)
	}
	if retVehicle.ID != retVehicle2.ID {
		t.Error("Vehicle is not the same")
	}
}

func TestCreateVehicleAlreadyExists(t *testing.T) {
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
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	_, err = CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	_, err = CreateVehicle(&Vehicle)
	if err == nil {
		t.Error("Vehicle was created")
	}
}

func TestGetVehicle(t *testing.T) {
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
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	retVehicle, err := CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	retVehicle2, err := GetVehicle(retVehicle.ID)
	if err != nil {
		t.Error(err)
	}
	if retVehicle.ID != retVehicle2.ID {
		t.Error("Vehicle is not the same")
	}
}

func TestGetVehicleNotFound(t *testing.T) {
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
	_, err = GetVehicle(0)
	if err == nil {
		t.Error("Vehicle was found")
	}
}

func TestGetVehicles(t *testing.T) {
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
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	retVehicle, err := CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	retVehicles, err := GetVehicles()
	if err != nil {
		t.Error(err)
	}
	if len(retVehicles) == 0 {
		t.Error("No vehicles found")
	}
	found := false
	for _, v := range retVehicles {
		if v.ID == retVehicle.ID {
			found = true
		}
	}
	if !found {
		t.Error("Vehicle not found")
	}
}

func TestDeleteVehicle(t *testing.T) {
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
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	retVehicle, err := CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	err = DeleteVehicle(retVehicle.ID)
	if err != nil {
		t.Error(err)
	}
	_, err = GetVehicle(retVehicle.ID)
	if err == nil {
		t.Error("Vehicle was found")
	}
}

func TestDeleteVehicleNotFound(t *testing.T) {
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
	err = DeleteVehicle(0)
	if err == nil {
		t.Error("Vehicle was found")
	}
}

func TestUpdateVehicle(t *testing.T) {
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
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	retVehicle, err := CreateVehicle(&Vehicle)
	if err != nil {
		t.Error(err)
	}
	retVehicle.Plate = "AA-321-AA"
	err = UpdateVehicle(retVehicle)
	if err != nil {
		t.Error(err)
	}
	retVehicle2, err := GetVehicle(retVehicle.ID)
	if err != nil {
		t.Error(err)
	}
	if retVehicle.Plate != retVehicle2.Plate {
		t.Error("Vehicle was not updated")
	}
}

func TestUpdateVehicleNotFound(t *testing.T) {
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
	var Vehicle = models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	}
	err = UpdateVehicle(&Vehicle)
	if err == nil {
		t.Error("Vehicle was found")
	}
}

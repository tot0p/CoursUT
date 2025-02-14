package parkingSpaceInformation

import (
	"context"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpace"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/models"
)

func setupDatabase(t *testing.T) func() {
	err := database.InitDatabase()
	require.NoError(t, err)
	return func() {
		err := database.CloseDatabase()
		require.NoError(t, err)
	}
}

func TestCreateParkingSpaceInformation(t *testing.T) {
	defer setupDatabase(t)()

	p1 := models.ParkingSpace{SpaceNumber: "A32"}
	retParkingSpace, err := parkingSpace.CreateParkingSpace(&p1)
	require.NoError(t, err)

	v := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	retVehicle, err := vehicle.CreateVehicle(&v)
	require.NoError(t, err)

	parkingSpaceInformation := models.ParkingSpaceInformation{
		ParkingSpaceID:  retParkingSpace.ID,
		VehicleID:       retVehicle.ID,
		ArrivalTime:     time.Now(),
		DepartureTime:   time.Now().Add(time.Hour),
		ParkingDuration: time.Hour,
	}
	retParkingSpaceInformation, err := CreateParkingSpaceInformation(&parkingSpaceInformation)
	require.NoError(t, err)
	assert.NotZero(t, retParkingSpaceInformation.ID)

	var p models.ParkingSpaceInformation
	err = database.Conn.QueryRowContext(context.Background(), "SELECT * FROM parking_space_information WHERE id = ?;", retParkingSpaceInformation.ID).Scan(&p.ID, &p.ParkingSpaceID, &p.VehicleID, &p.ArrivalTime, &p.DepartureTime, &p.ParkingDuration)
	require.NoError(t, err)
	assert.Equal(t, retParkingSpaceInformation.ID, p.ID)
}

func TestGetParkingSpaceInformation(t *testing.T) {
	defer setupDatabase(t)()

	p := models.ParkingSpace{SpaceNumber: "A32"}
	retParkingSpace, err := parkingSpace.CreateParkingSpace(&p)
	require.NoError(t, err)

	v := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	retVehicle, err := vehicle.CreateVehicle(&v)
	require.NoError(t, err)

	parkingSpaceInformation := models.ParkingSpaceInformation{
		ParkingSpaceID:  retParkingSpace.ID,
		VehicleID:       retVehicle.ID,
		ArrivalTime:     time.Now(),
		DepartureTime:   time.Now().Add(time.Hour),
		ParkingDuration: time.Hour,
	}
	retParkingSpaceInformation, err := CreateParkingSpaceInformation(&parkingSpaceInformation)
	require.NoError(t, err)

	retParkingSpaceInformation2, err := GetParkingSpaceInformation(retParkingSpaceInformation.ID)
	require.NoError(t, err)
	assert.Equal(t, retParkingSpaceInformation.ID, retParkingSpaceInformation2.ID)
}

func TestGetParkingSpaceInformationNotFound(t *testing.T) {
	defer setupDatabase(t)()

	_, err := GetParkingSpaceInformation(0)
	assert.Error(t, err)
}

func TestGetParkingSpaceInformations(t *testing.T) {
	defer setupDatabase(t)()

	p := models.ParkingSpace{SpaceNumber: "A32"}
	retParkingSpace, err := parkingSpace.CreateParkingSpace(&p)
	require.NoError(t, err)

	v := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	retVehicle, err := vehicle.CreateVehicle(&v)
	require.NoError(t, err)

	parkingSpaceInformation := models.ParkingSpaceInformation{
		ParkingSpaceID:  retParkingSpace.ID,
		VehicleID:       retVehicle.ID,
		ArrivalTime:     time.Now(),
		DepartureTime:   time.Now().Add(time.Hour),
		ParkingDuration: time.Hour,
	}
	_, err = CreateParkingSpaceInformation(&parkingSpaceInformation)
	require.NoError(t, err)

	retParkingSpaceInformations, err := GetParkingSpaceInformations()
	require.NoError(t, err)
	assert.NotEmpty(t, retParkingSpaceInformations)
}

func TestDeleteParkingSpaceInformation(t *testing.T) {
	defer setupDatabase(t)()

	p := models.ParkingSpace{SpaceNumber: "A32"}
	retParkingSpace, err := parkingSpace.CreateParkingSpace(&p)
	require.NoError(t, err)

	v := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	retVehicle, err := vehicle.CreateVehicle(&v)
	require.NoError(t, err)

	parkingSpaceInformation := models.ParkingSpaceInformation{
		ParkingSpaceID:  retParkingSpace.ID,
		VehicleID:       retVehicle.ID,
		ArrivalTime:     time.Now(),
		DepartureTime:   time.Now().Add(time.Hour),
		ParkingDuration: time.Hour,
	}
	retParkingSpaceInformation, err := CreateParkingSpaceInformation(&parkingSpaceInformation)
	require.NoError(t, err)

	err = DeleteParkingSpaceInformation(retParkingSpaceInformation.ID)
	require.NoError(t, err)

	_, err = GetParkingSpaceInformation(retParkingSpaceInformation.ID)
	assert.Error(t, err)
}

func TestDeleteParkingSpaceInformationNotFound(t *testing.T) {
	defer setupDatabase(t)()

	err := DeleteParkingSpaceInformation(0)
	assert.Error(t, err)
}

func TestUpdateParkingSpaceInformation(t *testing.T) {
	defer setupDatabase(t)()

	p := models.ParkingSpace{SpaceNumber: "A32"}
	retParkingSpace, err := parkingSpace.CreateParkingSpace(&p)
	require.NoError(t, err)

	v := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	retVehicle, err := vehicle.CreateVehicle(&v)
	require.NoError(t, err)

	parkingSpaceInformation := models.ParkingSpaceInformation{
		ParkingSpaceID:  retParkingSpace.ID,
		VehicleID:       retVehicle.ID,
		ArrivalTime:     time.Now(),
		DepartureTime:   time.Now().Add(time.Hour),
		ParkingDuration: time.Hour,
	}
	retParkingSpaceInformation, err := CreateParkingSpaceInformation(&parkingSpaceInformation)
	require.NoError(t, err)

	retParkingSpaceInformation.ArrivalTime = time.Now().Add(time.Hour)
	retParkingSpaceInformation.DepartureTime = time.Now().Add(2 * time.Hour)
	retParkingSpaceInformation.ParkingDuration = 2 * time.Hour
	err = UpdateParkingSpaceInformation(retParkingSpaceInformation)
	require.NoError(t, err)

	retParkingSpaceInformation2, err := GetParkingSpaceInformation(retParkingSpaceInformation.ID)
	require.NoError(t, err)
	assert.Equal(t, retParkingSpaceInformation.ID, retParkingSpaceInformation2.ID)
	assert.Equal(t, retParkingSpaceInformation.ArrivalTime.Unix(), retParkingSpaceInformation2.ArrivalTime.Unix())
	assert.Equal(t, retParkingSpaceInformation.DepartureTime.Unix(), retParkingSpaceInformation2.DepartureTime.Unix())
	assert.Equal(t, retParkingSpaceInformation.ParkingDuration, retParkingSpaceInformation2.ParkingDuration)
}

func TestUpdateParkingSpaceInformationNotFound(t *testing.T) {
	defer setupDatabase(t)()

	err := UpdateParkingSpaceInformation(&models.ParkingSpaceInformation{})
	assert.Error(t, err)
}

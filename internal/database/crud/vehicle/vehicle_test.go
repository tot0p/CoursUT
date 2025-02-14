package vehicle

import (
	"testing"

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

func TestCreateVehicle(t *testing.T) {
	defer setupDatabase(t)()

	vehicle := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	retVehicle, err := CreateVehicle(&vehicle)
	require.NoError(t, err)

	retVehicle2, err := GetVehicle(retVehicle.ID)
	require.NoError(t, err)
	assert.Equal(t, retVehicle.ID, retVehicle2.ID)
}

func TestCreateVehicleAlreadyExists(t *testing.T) {
	defer setupDatabase(t)()

	vehicle := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	_, err := CreateVehicle(&vehicle)
	require.NoError(t, err)

	_, err = CreateVehicle(&vehicle)
	assert.Error(t, err)
}

func TestGetVehicle(t *testing.T) {
	defer setupDatabase(t)()

	vehicle := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	retVehicle, err := CreateVehicle(&vehicle)
	require.NoError(t, err)

	retVehicle2, err := GetVehicle(retVehicle.ID)
	require.NoError(t, err)
	assert.Equal(t, retVehicle.ID, retVehicle2.ID)
}

func TestGetVehicleNotFound(t *testing.T) {
	defer setupDatabase(t)()

	_, err := GetVehicle(0)
	assert.Error(t, err)
}

func TestGetVehicles(t *testing.T) {
	defer setupDatabase(t)()

	vehicle := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	retVehicle, err := CreateVehicle(&vehicle)
	require.NoError(t, err)

	retVehicles, err := GetVehicles()
	require.NoError(t, err)
	assert.NotEmpty(t, retVehicles)

	found := false
	for _, v := range retVehicles {
		if v.ID == retVehicle.ID {
			found = true
		}
	}
	assert.True(t, found)
}

func TestDeleteVehicle(t *testing.T) {
	defer setupDatabase(t)()

	vehicle := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	retVehicle, err := CreateVehicle(&vehicle)
	require.NoError(t, err)

	err = DeleteVehicle(retVehicle.ID)
	require.NoError(t, err)

	_, err = GetVehicle(retVehicle.ID)
	assert.Error(t, err)
}

func TestDeleteVehicleNotFound(t *testing.T) {
	defer setupDatabase(t)()

	err := DeleteVehicle(0)
	assert.Error(t, err)
}

func TestUpdateVehicle(t *testing.T) {
	defer setupDatabase(t)()

	vehicle := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	retVehicle, err := CreateVehicle(&vehicle)
	require.NoError(t, err)

	retVehicle.Plate = "AA-321-AA"
	err = UpdateVehicle(retVehicle)
	require.NoError(t, err)

	retVehicle2, err := GetVehicle(retVehicle.ID)
	require.NoError(t, err)
	assert.Equal(t, retVehicle.Plate, retVehicle2.Plate)
}

func TestUpdateVehicleNotFound(t *testing.T) {
	defer setupDatabase(t)()

	vehicle := models.Vehicle{Plate: "AA-123-AA", VehicleType: models.Car}
	err := UpdateVehicle(&vehicle)
	assert.Error(t, err)
}

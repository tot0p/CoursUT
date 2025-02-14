package parkingSpace

import (
	"context"
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

func TestCreateParkingSpace(t *testing.T) {
	defer setupDatabase(t)()

	parkingSpace := models.ParkingSpace{SpaceNumber: "A32"}
	_, err := CreateParkingSpace(&parkingSpace)
	require.NoError(t, err)
	assert.NotEmpty(t, parkingSpace.SpaceNumber)

	var p models.ParkingSpace
	err = database.Conn.QueryRowContext(context.Background(), "SELECT * FROM parking_space WHERE id = ?;", parkingSpace.ID).Scan(&p.ID, &p.SpaceNumber)
	require.NoError(t, err)
	assert.Equal(t, parkingSpace, p)
}

func TestCreateParkingSpaceAlreadyExists(t *testing.T) {
	defer setupDatabase(t)()

	parkingSpace := models.ParkingSpace{SpaceNumber: "A32"}
	_, err := CreateParkingSpace(&parkingSpace)
	require.NoError(t, err)

	_, err = CreateParkingSpace(&parkingSpace)
	assert.Error(t, err)
}

func TestGetParkingSpace(t *testing.T) {
	defer setupDatabase(t)()

	parkingSpace := models.ParkingSpace{SpaceNumber: "A32"}
	temp, err := CreateParkingSpace(&parkingSpace)
	require.NoError(t, err)
	parkingSpace = *temp

	p, err := GetParkingSpace(parkingSpace.ID)
	require.NoError(t, err)
	assert.Equal(t, parkingSpace, *p)
}

func TestGetParkingSpaceNotFound(t *testing.T) {
	defer setupDatabase(t)()

	_, err := GetParkingSpace(1)
	assert.Error(t, err)
}

func TestGetParkingSpaces(t *testing.T) {
	defer setupDatabase(t)()

	parkingSpace := models.ParkingSpace{SpaceNumber: "A32"}
	_, err := CreateParkingSpace(&parkingSpace)
	require.NoError(t, err)

	parkingSpaces, err := GetParkingSpaces()
	require.NoError(t, err)
	assert.NotEmpty(t, parkingSpaces)
}

func TestDeleteParkingSpace(t *testing.T) {
	defer setupDatabase(t)()

	parkingSpace := models.ParkingSpace{SpaceNumber: "A32"}
	_, err := CreateParkingSpace(&parkingSpace)
	require.NoError(t, err)

	err = DeleteParkingSpace(parkingSpace.ID)
	require.NoError(t, err)

	_, err = GetParkingSpace(parkingSpace.ID)
	assert.Error(t, err)
}

func TestDeleteParkingSpaceNotFound(t *testing.T) {
	defer setupDatabase(t)()

	err := DeleteParkingSpace(1)
	assert.Error(t, err)
}

func TestUpdateParkingSpace(t *testing.T) {
	defer setupDatabase(t)()

	parkingSpace := models.ParkingSpace{SpaceNumber: "A32"}
	temp, err := CreateParkingSpace(&parkingSpace)
	require.NoError(t, err)
	parkingSpace.ID = temp.ID
	parkingSpace.SpaceNumber = "A33"

	err = UpdateParkingSpace(&parkingSpace)
	require.NoError(t, err)

	p, err := GetParkingSpace(parkingSpace.ID)
	require.NoError(t, err)
	assert.Equal(t, "A33", p.SpaceNumber)
}

func TestUpdateParkingSpaceNotFound(t *testing.T) {
	defer setupDatabase(t)()

	parkingSpace := models.ParkingSpace{ID: 1, SpaceNumber: "A32"}
	err := UpdateParkingSpace(&parkingSpace)
	assert.Error(t, err)
}

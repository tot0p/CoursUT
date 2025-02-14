package database_test

import (
	"github.com/tot0p/CoursUT/internal/database"
	"testing"

	_ "github.com/marcboeker/go-duckdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitDatabase(t *testing.T) {
	err := database.InitDatabase()
	require.NoError(t, err)
	assert.NotNil(t, database.Conn)
}

func TestCloseDatabase(t *testing.T) {
	err := database.InitDatabase()
	require.NoError(t, err)

	err = database.CloseDatabase()
	assert.NoError(t, err)
}

func TestCreateTableVehicle(t *testing.T) {
	err := database.InitDatabase()
	require.NoError(t, err)

	err = database.CreateTableVehicle()
	assert.NoError(t, err)
}

func TestCreateTableParkingSpace(t *testing.T) {
	err := database.InitDatabase()
	require.NoError(t, err)

	err = database.CreateTableParkingSpace()
	assert.NoError(t, err)
}

func TestCreateTableParkingSpaceInformation(t *testing.T) {
	err := database.InitDatabase()
	require.NoError(t, err)

	err = database.CreateTableParkingSpaceInformation()
	assert.NoError(t, err)
}

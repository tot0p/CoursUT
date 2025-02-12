package vehicle

import (
	"context"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/models"
)

// CreateVehicle creates a vehicle in the database
func CreateVehicle(vehicle *models.Vehicle) (*models.Vehicle, error) {
	res, err := database.Conn.ExecContext(context.Background(), "INSERT INTO vehicle (plate, vehicle_type) VALUES (?, ?);", vehicle.Plate, vehicle.VehicleType)
	if err != nil {
		return nil, err
	}
	tempId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	vehicle.ID = int(tempId)
	return vehicle, nil
}

// GetVehicle returns a vehicle from the database
func GetVehicle(id int) (*models.Vehicle, error) {
	vehicle := models.Vehicle{}
	err := database.Conn.QueryRowContext(context.Background(), "SELECT * FROM vehicle WHERE id = ?;", id).Scan(&vehicle.ID, &vehicle.Plate, &vehicle.VehicleType)
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// GetVehicles returns all vehicles from the database
func GetVehicles() ([]models.Vehicle, error) {
	rows, err := database.Conn.QueryContext(context.Background(), "SELECT * FROM vehicle;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var vehicles []models.Vehicle
	for rows.Next() {
		vehicle := models.Vehicle{}
		err = rows.Scan(&vehicle.ID, &vehicle.Plate, &vehicle.VehicleType)
		if err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}

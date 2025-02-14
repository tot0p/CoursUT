package vehicle

import (
	"context"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/models"
)

// CreateVehicle creates a vehicle in the database
func CreateVehicle(vehicle *models.Vehicle) (*models.Vehicle, error) {
	res := database.Conn.QueryRowContext(context.Background(), "INSERT INTO vehicle (plate, vehicle_type) VALUES (?, ?) RETURNING id;", vehicle.Plate, vehicle.VehicleType)
	err := res.Scan(&vehicle.ID)
	if err != nil {
		return nil, err
	}
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

func DeleteVehicle(id int) error {
	_, err := database.Conn.ExecContext(context.Background(), "DELETE FROM vehicle WHERE id = ?;", id)
	return err
}

func UpdateVehicle(vehicle *models.Vehicle) error {
	_, err := database.Conn.ExecContext(context.Background(), "UPDATE vehicle SET plate = ?, vehicle_type = ? WHERE id = ?;", vehicle.Plate, vehicle.VehicleType, vehicle.ID)
	return err
}

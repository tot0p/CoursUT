package vehicle

import (
	"context"
	"fmt"
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
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning rows: %w", err)
	}
	return vehicles, nil
}

func DeleteVehicle(id int) error {
	res, err := database.Conn.ExecContext(context.Background(), "DELETE FROM vehicle WHERE id = ?;", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("vehicle with id %d not found", id)
	}
	return nil
}

func UpdateVehicle(vehicle *models.Vehicle) error {
	// to Update Use a Delete ... Return ... and a Create
	res, err := database.Conn.ExecContext(context.Background(), "DELETE FROM vehicle WHERE id = ?;", vehicle.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("vehicle with id %d not found", vehicle.ID)
	}
	res2 := database.Conn.QueryRowContext(context.Background(), "INSERT INTO vehicle (id,plate, vehicle_type) VALUES (?,?, ?) RETURNING id;", vehicle.ID, vehicle.Plate, vehicle.VehicleType)
	err = res2.Scan(&vehicle.ID)
	if err != nil {
		return err
	}
	return nil
}

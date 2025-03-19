package parkingSpaceInformation

import (
	"context"
	"fmt"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/models"
)

func CreateParkingSpaceInformation(information *models.ParkingSpaceInformation) (*models.ParkingSpaceInformation, error) {
	information.ArrivalTime = information.ArrivalTime.UTC()
	information.DepartureTime = information.DepartureTime.UTC()
	res := database.Conn.QueryRowContext(context.Background(), "INSERT INTO parking_space_information (parking_space_id, vehicle_id, arrival_time, parking_duration) VALUES (?, ?, ?, ?) RETURNING id;", information.ParkingSpaceID, information.VehicleID, information.ArrivalTime, information.ParkingDuration)
	err := res.Scan(&information.ID)
	if err != nil {
		return nil, err
	}
	return information, nil
}

func GetParkingSpaceInformation(id int) (*models.ParkingSpaceInformation, error) {
	information := models.ParkingSpaceInformation{}
	err := database.Conn.QueryRowContext(context.Background(), "SELECT * FROM parking_space_information WHERE id = ?;", id).Scan(&information.ID, &information.VehicleID, &information.ParkingSpaceID, &information.ArrivalTime, &information.DepartureTime, &information.ParkingDuration)
	if err != nil {
		return nil, err
	}
	return &information, nil
}

func GetParkingSpaceInformations() ([]models.ParkingSpaceInformation, error) {
	rows, err := database.Conn.QueryContext(context.Background(), "SELECT * FROM parking_space_information;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var informations []models.ParkingSpaceInformation
	for rows.Next() {
		information := models.ParkingSpaceInformation{}
		err = rows.Scan(&information.ID, &information.VehicleID, &information.ParkingSpaceID, &information.ArrivalTime, &information.DepartureTime, &information.ParkingDuration)
		if err != nil {
			return nil, err
		}
		informations = append(informations, information)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning rows: %w", err)
	}
	return informations, nil
}

func DeleteParkingSpaceInformation(id int) error {
	res, err := database.Conn.ExecContext(context.Background(), "DELETE FROM parking_space_information WHERE id = ?;", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("parking space information with id %d not found", id)
	}
	return nil
}

func UpdateParkingSpaceInformation(information *models.ParkingSpaceInformation) error {
	// to Update Use a Delete ... Return ... and a Create
	information.DepartureTime = information.DepartureTime.UTC()
	information.ArrivalTime = information.ArrivalTime.UTC()
	res, err := database.Conn.ExecContext(context.Background(), "DELETE FROM parking_space_information WHERE id = ?;", information.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("parking space information with id %d not found", information.ID)
	}
	res2 := database.Conn.QueryRowContext(context.Background(), "INSERT INTO parking_space_information (id, parking_space_id, vehicle_id, arrival_time, departure_time, parking_duration) VALUES (?, ?, ?, ?, ?, ?) RETURNING id;", information.ID, information.ParkingSpaceID, information.VehicleID, information.ArrivalTime, information.DepartureTime, information.ParkingDuration)
	err = res2.Scan(&information.ID)
	if err != nil {
		return err
	}
	return nil
}

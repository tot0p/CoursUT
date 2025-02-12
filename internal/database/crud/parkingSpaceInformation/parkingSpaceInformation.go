package parkingSpaceInformation

import (
	"context"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/models"
	"time"
)

func CreateParkingSpaceInformation(information *models.ParkingSpaceInformation) (*models.ParkingSpaceInformation, error) {
	information.ArrivalTime = time.Now()
	res, err := database.Conn.ExecContext(context.Background(), "INSERT INTO parking_space_information (parking_space_id, vehicle_id, arrival_time, parking_duration) VALUES (?, ?, ?, ?);", information.ParkingSpaceID, information.VehicleID, information.ArrivalTime, information.ParkingDuration)
	if err != nil {
		return nil, err
	}
	tempId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	information.ID = int(tempId)
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
	return informations, nil
}

func DeleteParkingSpaceInformation(id int) error {
	_, err := database.Conn.ExecContext(context.Background(), "DELETE FROM parking_space_information WHERE id = ?;", id)
	return err
}

func UpdateParkingSpaceInformation(information *models.ParkingSpaceInformation) error {
	_, err := database.Conn.ExecContext(context.Background(), "UPDATE parking_space_information SET vehicle_id = ?, parking_space_id = ?, arrival_time = ?, departure_time = ?, parking_duration = ? WHERE id = ?;", information.VehicleID, information.ParkingSpaceID, information.ArrivalTime, information.DepartureTime, information.ParkingDuration, information.ID)
	return err
}

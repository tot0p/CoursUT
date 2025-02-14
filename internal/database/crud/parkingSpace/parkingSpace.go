package parkingSpace

import (
	"context"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/models"
)

func CreateParkingSpace(parkingSpace *models.ParkingSpace) (*models.ParkingSpace, error) {
	res := database.Conn.QueryRowContext(context.Background(), "INSERT INTO parking_space (space_number) VALUES (?) RETURNING id;", parkingSpace.SpaceNumber)
	err := res.Scan(&parkingSpace.ID)
	if err != nil {
		return nil, err
	}
	return parkingSpace, nil
}

func GetParkingSpace(id int) (*models.ParkingSpace, error) {
	parkingSpace := models.ParkingSpace{}
	err := database.Conn.QueryRowContext(context.Background(), "SELECT * FROM parking_space WHERE id = ?;", id).Scan(&parkingSpace.ID, &parkingSpace.SpaceNumber)
	if err != nil {
		return nil, err
	}
	return &parkingSpace, nil
}

func GetParkingSpaces() ([]models.ParkingSpace, error) {
	rows, err := database.Conn.QueryContext(context.Background(), "SELECT * FROM parking_space;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var parkingSpaces []models.ParkingSpace
	for rows.Next() {
		parkingSpace := models.ParkingSpace{}
		err = rows.Scan(&parkingSpace.ID, &parkingSpace.SpaceNumber)
		if err != nil {
			return nil, err
		}
		parkingSpaces = append(parkingSpaces, parkingSpace)
	}
	return parkingSpaces, nil
}

func DeleteParkingSpace(id int) error {
	_, err := database.Conn.ExecContext(context.Background(), "DELETE FROM parking_space WHERE id = ?;", id)
	return err
}

func UpdateParkingSpace(parkingSpace *models.ParkingSpace) error {
	_, err := database.Conn.ExecContext(context.Background(), "UPDATE parking_space SET space_number = ? WHERE id = ?;", parkingSpace.SpaceNumber, parkingSpace.ID)
	return err
}

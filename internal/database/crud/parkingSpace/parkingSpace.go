package parkingSpace

import (
	"context"
	"fmt"
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
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning rows: %w", err)
	}
	return parkingSpaces, nil
}

func DeleteParkingSpace(id int) error {
	res, err := database.Conn.ExecContext(context.Background(), "DELETE FROM parking_space WHERE id = ?;", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("parking space with id %d not found", id)
	}
	return nil
}

func UpdateParkingSpace(parkingSpace *models.ParkingSpace) error {
	// to Update Use a Delete ... Return ... and a Create
	res, err := database.Conn.ExecContext(context.Background(), "DELETE FROM parking_space WHERE id = ?;", parkingSpace.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("parking space with id %d not found", parkingSpace.ID)
	}
	res2 := database.Conn.QueryRowContext(context.Background(), "INSERT INTO parking_space (id, space_number) VALUES (?, ?) RETURNING id;", parkingSpace.ID, parkingSpace.SpaceNumber)
	err = res2.Scan(&parkingSpace.ID)
	if err != nil {
		return err
	}
	return nil
}

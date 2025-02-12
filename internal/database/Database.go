package database

import (
	"context"
	"database/sql"
	_ "github.com/marcboeker/go-duckdb"
)

type dbConn struct {
	*sql.Conn
}

var Conn dbConn

// InitDatabase initializes the database
func InitDatabase() error {
	var err error
	db, err := sql.Open("duckdb", "")
	if err != nil {
		return err
	}
	Conn.Conn, err = db.Conn(context.Background())
	if err != nil {
		return err
	}
	err = initSchema()
	if err != nil {
		return err
	}
	return nil
}

// InitDatabaseFromFilename creates or opens a database file
func InitDatabaseFromFilename(filename string) error {
	var err error
	db, err := sql.Open("duckdb", filename)
	if err != nil {
		return err
	}
	Conn.Conn, err = db.Conn(context.Background())
	if err != nil {
		return err
	}
	err = initSchema()
	if err != nil {
		return err
	}
	return nil
}

func initSchema() error {
	transac, err := Conn.BeginTx(context.Background(), nil)
	if err != nil {
		err1 := transac.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}
	err = CreateTableVehicle()
	if err != nil {
		err1 := transac.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}
	err = CreateTableParkingSpace()
	if err != nil {
		err1 := transac.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}
	err = CreateTableParkingSpaceInformation()
	if err != nil {
		err1 := transac.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}
	err = transac.Commit()
	if err != nil {
		err1 := transac.Rollback()
		if err1 != nil {
			return err1
		}
		return err
	}
	return nil
}

func CreateTableParkingSpace() error {
	query := `
	CREATE SEQUENCE id_parking_space START 1;
	CREATE TABLE IF NOT EXISTS parking_space (
		id INT PRIMARY KEY DEFAULT NEXTVAL('id_parking_space'),
		space_number TEXT UNIQUE NOT NULL
	);`
	_, err := Conn.ExecContext(context.Background(), query)
	return err
}

func CreateTableParkingSpaceInformation() error {
	query := `
	CREATE SEQUENCE id_parking_space_information START 1;
	CREATE TABLE IF NOT EXISTS parking_space_information (
		id INT PRIMARY KEY DEFAULT nextval('id_parking_space_information'),
		vehicle_id INT NOT NULL,
		parking_space_id INT NOT NULL,
		arrival_time TIMESTAMP NOT NULL,
		departure_time TIMESTAMP NOT NULL,
		parking_duration INTERVAL NOT NULL,
        CONSTRAINT fk_vehicle_id FOREIGN KEY (vehicle_id) REFERENCES vehicle(id),
        CONSTRAINT fk_parking_space_id FOREIGN KEY (parking_space_id) REFERENCES parking_space(id)
	);`
	_, err := Conn.ExecContext(context.Background(), query)
	return err
}

func CreateTableVehicle() error {
	query := `
	CREATE SEQUENCE id_vehicle START 1;
	CREATE TABLE IF NOT EXISTS vehicle (
		id INT PRIMARY KEY DEFAULT NEXTVAL('id_vehicle'),
		plate TEXT NOT NULL,
		vehicle_type INT NOT NULL
	);`
	_, err := Conn.ExecContext(context.Background(), query)
	return err
}

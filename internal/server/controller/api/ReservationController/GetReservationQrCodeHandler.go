package ReservationController

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpaceInformation"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"strconv"
)

type myBuffer struct {
	bytes.Buffer
}

func (w *myBuffer) Write(p []byte) (n int, err error) {
	return w.Buffer.Write(p)
}

func (w *myBuffer) Close() error {
	return nil
}

// GetReservationQrCodeHandler generates a QR code for a reservation
func GetReservationQrCodeHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	reservationID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid reservation ID"})
	}

	reservation, err := parkingSpaceInformation.GetParkingSpaceInformation(reservationID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reservation not found"})
	}

	var data = make(map[string]interface{})
	data["id"] = reservation.ID
	data["parkingSpaceID"] = reservation.ParkingSpaceID
	data["vehicleID"] = reservation.VehicleID

	rawData, err := json.Marshal(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate QR code"})
	}

	qrCode, err := qrcode.New(string(rawData))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate QR code"})
	}

	var buf myBuffer
	writer := standard.NewWithWriter(&buf)
	if err := qrCode.Save(writer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to write QR code to buffer"})
	}
	//set the content type to image/png
	c.Set(fiber.HeaderContentType, "image/png")
	return c.Status(fiber.StatusOK).Send(buf.Bytes())
}

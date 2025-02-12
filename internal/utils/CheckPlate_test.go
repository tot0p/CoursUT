package utils

import "testing"

func TestCheckPlate(t *testing.T) {
	plaques := []struct {
		plate  string
		result bool
	}{
		{
			plate:  "AA-123-AA", // Nouveau format
			result: true,
		},
		{
			plate:  "1234 ABC 75", // Ancien format
			result: true,
		},
		{
			plate:  "123 AB 01", // Ancien format
			result: true,
		},
		{
			plate:  "AB-123-C", // Invalide
			result: false,
		},
		{
			plate:  "1234 2A 75", // Invalide
			result: false,
		},
		{
			plate:  "test",
			result: false,
		},
	}

	for _, plaque := range plaques {
		if CheckPlate(plaque.plate) != plaque.result {
			t.Errorf("CheckPlate(%s) = %t; want %t", plaque.plate, !plaque.result, plaque.result)
		}
	}
}

package models

import "testing"

func TestIsValidVehicleType(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{
			name:  "Unknown",
			value: 0,
			want:  true,
		},
		{
			name:  "Car",
			value: 1,
			want:  true,
		},
		{
			name:  "Truck",
			value: 2,
			want:  true,
		},
		{
			name:  "Invalid",
			value: 3,
			want:  false,
		},
		{
			name:  "Invalid",
			value: -1,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidVehicleType(tt.value); got != tt.want {
				t.Errorf("IsValidVehicleType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidParkingSpace(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{
			name:  "Valid",
			value: "A1",
			want:  true,
		},
		{
			name:  "Valid",
			value: "B23",
			want:  true,
		},
		{
			name:  "Valid",
			value: "C999",
			want:  true,
		},
		{
			name:  "Invalid",
			value: "D1234",
			want:  false,
		},
		{
			name:  "Invalid",
			value: "E12345",
			want:  false,
		},
		{
			name:  "Invalid",
			value: "F123456",
			want:  false,
		},
		{
			name:  "Invalid",
			value: "G",
			want:  false,
		},
		{
			name:  "Valid",
			value: "H123",
			want:  true,
		},
		{
			name:  "Invalid",
			value: "I1234",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidParkingSpace(tt.value); got != tt.want {
				t.Errorf("IsValidParkingSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

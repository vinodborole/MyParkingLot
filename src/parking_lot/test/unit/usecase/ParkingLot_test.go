package usecase

import (
	"os"
	"parking_lot/gateway"
	"parking_lot/usecase"
	"testing"
)

func TestMain(m *testing.M) {
	retCode := m.Run()
	os.Exit(retCode)
}

const (
	WrongSizeParkingLotError = "parking Lot of Size <= 0 cannot be created"
)

func TestParkLotInitialize_Success(t *testing.T) {
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	actual := pLInteract.ParkingLot.Initialize(6)
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", "No Error", actual)
	}
}
func TestParkLotInitialize_Failure(t *testing.T) {
	expected := WrongSizeParkingLotError
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	actual := pLInteract.ParkingLot.Initialize(-1).Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", expected, actual)
	}
}

func TestParkLotFull(t *testing.T) {
	expected := false
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	pLInteract.ParkingLot.Initialize(5)
	actual, _ := pLInteract.ParkingLot.IsParkingLotFull()
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got error:  '%v'", expected, actual)
	}
}

package usecase

import (
	"errors"
	"os"
	"parking_lot/domain"
	"parking_lot/gateway"
	"parking_lot/test/unit/mock"
	"parking_lot/usecase"
	"testing"
)

func TestMain(m *testing.M) {
	retCode := m.Run()
	os.Exit(retCode)
}

const (
	//WrongSizeParkingLotError wrong size parking lot
	WrongSizeParkingLotError = "parking Lot of Size <= 0 cannot be created"
	//ParkingLotFullError error text for parking lot full
	ParkingLotFullError = "sorry, parking lot is full"
	//ParkingLotNotCreatedError error text for parking lot not created error
	ParkingLotNotCreatedError = "parking Lot not created"
	//NotFoundError not found error
	NotFoundError = "not found"
)

func TestParkingLotInitialize_Success(t *testing.T) {
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	actual := pLInteract.ParkingLot.Initialize(6)
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", "No Error", actual)
	}
}
func TestParkingLotInitialize_Failure(t *testing.T) {
	expected := WrongSizeParkingLotError
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	actual := pLInteract.ParkingLot.Initialize(-1).Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", expected, actual)
	}
}

func TestParkingLotFull(t *testing.T) {
	expected := false
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	pLInteract.ParkingLot.Initialize(5)
	actual, _ := pLInteract.ParkingLot.IsParkingLotFull()
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got error:  '%v'", expected, actual)
	}
}

func TestParkingLotCreated(t *testing.T) {
	expected := true
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	pLInteract.ParkingLot.Initialize(5)
	actual, _ := pLInteract.ParkingLot.IsParkingLotCreated()
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got error:  '%v'", expected, actual)
	}
}

func TestIsParkingLotCreatedMock(t *testing.T) {
	expected := ParkingLotNotCreatedError
	MockParkingLotRepository := mock.ParkingLotRepository{MockIsParkingLotCreated: func() (b bool, err error) {
		return false, errors.New(ParkingLotNotCreatedError)
	}}
	interact := usecase.ParkingLotInteract{ParkingLot: &MockParkingLotRepository}
	car := domain.Create("MH-12-QY-1234", "White")
	actual := interact.Park(car).Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestParkingLotNotCreated(t *testing.T) {
	expected := false
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	actual, err := pLInteract.ParkingLot.IsParkingLotCreated()
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got error:  '%v'", expected, actual)
	}
	if err == nil {
		t.Errorf("Test failed, expected: '%v', got error:  '%v'", ParkingLotNotCreatedError, "nil")
	}
}
func TestParkSuccess(t *testing.T) {
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	pLInteract.ParkingLot.Initialize(6)
	car := domain.Create("MH-12-QY-1234", "White")
	actual := pLInteract.Park(car)
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", "No Error", actual)
	}
}

func TestParkingLotFullError(t *testing.T) {
	expected := ParkingLotFullError
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	pLInteract.ParkingLot.Initialize(1)
	car1 := domain.Create("MH-12-QY-1234", "White")
	pLInteract.Park(car1)
	car2 := domain.Create("MH-14-QY-1634", "Black")
	actual := pLInteract.Park(car2)
	if actual == nil {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", "No Error", actual)
	}
	if actual.Error() != expected {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", expected, actual.Error())
	}
}
func TestIsParkingLotFullMock(t *testing.T) {
	expected := ParkingLotFullError
	MockParkingLotRepository := mock.ParkingLotRepository{MockIsParkingLotFull: func() (b bool, err error) {
		return false, errors.New(ParkingLotFullError)
	}}
	interact := usecase.ParkingLotInteract{ParkingLot: &MockParkingLotRepository}
	car := domain.Create("MH-12-QY-1234", "White")
	actual := interact.Park(car).Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestParkSuccess_GetRegNo(t *testing.T) {
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	pLInteract.ParkingLot.Initialize(6)
	car1 := domain.Create("MH-12-QY-1234", "White")
	pLInteract.Park(car1)
	car2 := domain.Create("MH-13-AY-1534", "Black")
	pLInteract.Park(car2)
	car3 := domain.Create("MH-14-BY-1434", "Red")
	pLInteract.Park(car3)
	car4 := domain.Create("MH-15-CY-1254", "White")
	pLInteract.Park(car4)
	_, err := pLInteract.GetRegNosForCarsWithColor("White", true)
	if err != nil {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", "No Error", err.Error())
	}
	slot, _ := pLInteract.GetSlotNoForRegNo("MH-15-CY-1254", true)
	if slot != 4 {
		t.Errorf("Test failed, expected slot: '%d', got:  '%d'", 4, slot)
	}
	slots, err := pLInteract.GetSlotNosForCarsWithColor("White")
	if len(slots) != 2 {
		t.Errorf("Test failed, expected slot: '%d', got:  '%d'", 2, len(slots))
	}
	expectedSlot := func(expected int) bool {
		for _, slot := range slots {
			if slot == expected {
				return true
			}
		}
		return false
	}
	if !expectedSlot(4) {
		t.Errorf("Test failed, expected slot: '%d'", 4)
	}
	if !expectedSlot(1) {
		t.Errorf("Test failed, expected slot: '%d'", 2)
	}
}
func TestParkLeaveSuccess(t *testing.T) {
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	pLInteract.ParkingLot.Initialize(6)
	car := domain.Create("MH-12-QY-1234", "White")
	actual := pLInteract.Park(car)
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", "No Error", actual)
	}
	actual = pLInteract.Leave(1)
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", "No Error", actual)
	}
}

func TestParkSuccess_LeaveFailWithIncorrectSlot(t *testing.T) {
	parkingLotRepo := gateway.ParkingLotRepository{}
	pLInteract := usecase.ParkingLotInteract{ParkingLot: &parkingLotRepo}
	pLInteract.ParkingLot.Initialize(6)
	car := domain.Create("MH-12-QY-1234", "White")
	pLInteract.Park(car)
	expected := NotFoundError
	actual := pLInteract.Leave(2).Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got error:  '%s'", "No Error", actual)
	}
}

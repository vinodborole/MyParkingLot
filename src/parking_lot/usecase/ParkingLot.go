package usecase

import (
	"parking_lot/domain"
	Interact "parking_lot/usecase/interactorinterface"
)

//ParkingLotInteract parking lot interact
type ParkingLotInteract struct {
	ParkingLot Interact.ParkingLotRepository
}

//Park park car in parking lot
func (sh *ParkingLotInteract) Park(car domain.Car) error {

	return nil
}

//Leave leave car from parking lot
func (sh *ParkingLotInteract) Leave(slot int) error {
	return nil
}

//Status check status of parking lot
func (sh *ParkingLotInteract) Status() error {
	return nil
}

//GetRegNosForCarsWithColor get registration number for cars with given color
func (sh *ParkingLotInteract) GetRegNosForCarsWithColor(color string, print bool) ([]string, error) {
	return nil, nil
}

//GetSlotNosForCarsWithColor get slot nos for cars with given color
func (sh *ParkingLotInteract) GetSlotNosForCarsWithColor(color string) ([]int, error) {
	return nil, nil
}

//GetSlotNoForRegNo get slot nos for registration number
func (sh *ParkingLotInteract) GetSlotNoForRegNo(regNo string, print bool) (int, error) {
	return 0, nil
}

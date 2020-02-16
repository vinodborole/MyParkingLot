package mock

import (
	"parking_lot/domain"
	"parking_lot/helpers"
)

//ParkingLotRepository mock functions for parking lot
type ParkingLotRepository struct {
	MockInitialize                     func(numberOfSlots int) error
	MockVerifySlotInitialization       func(numberOfSlots int) (bool, error)
	MockIsParkingLotFull               func() (bool, error)
	MockIsParkingLotCreated            func() (bool, error)
	MockGetMaxParkingLotSize           func() int
	MockGetEmptySlots                  func() helpers.IntHeap
	MockPopEmptySlot                   func() interface{}
	MockPushEmptySlots                 func(slot int)
	MockMapRegNoToSlot                 func(regNo string, slot int)
	MockUnmapRegNo                     func(regNo string)
	MockMapSlotToCar                   func(slot int, car domain.Car)
	MockUnmapSlot                      func(slot int)
	MockMapColorToRegNo                func(color string, regNo string)
	MockUnmapRegNoFromColor            func(color string, regNo string)
	MockGetCarAtSlot                   func(slot int) (domain.Car, bool)
	MockGetSlotForRegistrationNumber   func(regNo string) (int, bool)
	MockGetRegistrationNumberOfByColor func(color string) (map[string]bool, bool)
}

//Initialize mock initialise parking lot
func (pL *ParkingLotRepository) Initialize(numberOfSlots int) error {
	if pL.MockInitialize != nil {
		return pL.MockInitialize(numberOfSlots)
	}
	return nil
}

//VerifySlotInitialization mock for verify slot initialization
func (pL *ParkingLotRepository) VerifySlotInitialization(numberOfSlots int) (bool, error) {
	if pL.MockVerifySlotInitialization != nil {
		return pL.MockVerifySlotInitialization(numberOfSlots)
	}
	return false, nil
}

//IsParkingLotFull mock for is parking lot full
func (pL *ParkingLotRepository) IsParkingLotFull() (bool, error) {
	if pL.MockIsParkingLotFull != nil {
		return pL.MockIsParkingLotFull()
	}
	return false, nil
}

//IsParkingLotCreated mock for is parking lot created
func (pL *ParkingLotRepository) IsParkingLotCreated() (bool, error) {
	if pL.MockIsParkingLotCreated != nil {
		return pL.MockIsParkingLotCreated()
	}
	return false, nil
}

//GetMaxParkingLotSize mock for get max parking lot size
func (pL *ParkingLotRepository) GetMaxParkingLotSize() int {
	if pL.MockGetMaxParkingLotSize != nil {
		return pL.MockGetMaxParkingLotSize()
	}
	return 0
}

//GetEmptySlots mock for get empty slots
func (pL *ParkingLotRepository) GetEmptySlots() helpers.IntHeap {
	if pL.MockGetEmptySlots != nil {
		return pL.MockGetEmptySlots()
	}
	return nil
}

//PopEmptySlot mock for pop empty slot
func (pL *ParkingLotRepository) PopEmptySlot() interface{} {
	if pL.MockPopEmptySlot != nil {
		return pL.MockPopEmptySlot()
	}
	return nil
}

//PushEmptySlots mock for push empty slots
func (pL *ParkingLotRepository) PushEmptySlots(slot int) {
	if pL.MockPushEmptySlots != nil {
		pL.MockPushEmptySlots(slot)
	}
}

//MapRegNoToSlot mock for mapping reg No to slot
func (pL *ParkingLotRepository) MapRegNoToSlot(regNo string, slot int) {
	if pL.MockMapRegNoToSlot != nil {
		pL.MockMapRegNoToSlot(regNo, slot)
	}
}

//UnmapRegNo mock for unmapping reg no
func (pL *ParkingLotRepository) UnmapRegNo(regNo string) {
	if pL.MockUnmapRegNo != nil {
		pL.MockUnmapRegNo(regNo)
	}
}

//MapSlotToCar mock for mapping slot to car
func (pL *ParkingLotRepository) MapSlotToCar(slot int, car domain.Car) {
	if pL.MockMapSlotToCar != nil {
		pL.MockMapSlotToCar(slot, car)
	}
}

//UnmapSlot mock to unmap slot
func (pL *ParkingLotRepository) UnmapSlot(slot int) {
	if pL.MockUnmapSlot != nil {
		pL.MockUnmapSlot(slot)
	}
}

//MapColorToRegNo mock to map color to reg no
func (pL *ParkingLotRepository) MapColorToRegNo(color string, regNo string) {
	if pL.MockMapColorToRegNo != nil {
		pL.MockMapColorToRegNo(color, regNo)
	}
}

//UnmapRegNoFromColor mock to unmap reg no from color
func (pL *ParkingLotRepository) UnmapRegNoFromColor(color string, regNo string) {
	if pL.MockUnmapRegNoFromColor != nil {
		pL.MockUnmapRegNoFromColor(color, regNo)
	}
}

//GetCarAtSlot mock for get car at slot
func (pL *ParkingLotRepository) GetCarAtSlot(slot int) (domain.Car, bool) {
	if pL.MockGetCarAtSlot != nil {
		return pL.MockGetCarAtSlot(slot)
	}
	return domain.Car{}, false
}

//GetSlotForRegistrationNumber mock for get slot for reg no
func (pL *ParkingLotRepository) GetSlotForRegistrationNumber(regNo string) (int, bool) {
	if pL.MockGetSlotForRegistrationNumber != nil {
		return pL.MockGetSlotForRegistrationNumber(regNo)
	}
	return 0, false
}

//GetRegistrationNumberOfByColor mock for get reg no by color
func (pL *ParkingLotRepository) GetRegistrationNumberOfByColor(color string) (map[string]bool, bool) {
	if pL.MockGetRegistrationNumberOfByColor != nil {
		return pL.MockGetRegistrationNumberOfByColor(color)
	}
	return map[string]bool{}, false
}

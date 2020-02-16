package interactorinterface

import (
	"parking_lot/domain"
	"parking_lot/helpers"
)

//ParkingLotRepository parking lot repository functions
type ParkingLotRepository interface {
	Initialize(numberOfSlots int) error
	VerifySlotInitialization(numberOfSlots int) (bool, error)
	IsParkingLotFull() (bool, error)
	IsParkingLotCreated() (bool, error)
	GetMaxParkingLotSize() int
	GetEmptySlots() helpers.IntHeap
	PopEmptySlot() interface{}
	PushEmptySlots(slot int)
	MapRegNoToSlot(regNo string, slot int)
	UnmapRegNo(regNo string)
	MapSlotToCar(slot int, car domain.Car)
	UnmapSlot(slot int)
	MapColorToRegNo(color string, regNo string)
	UnmapRegNoFromColor(color string, regNo string)
	GetCarAtSlot(slot int) (domain.Car, bool)
	GetSlotForRegistrationNumber(regNo string) (int, bool)
	GetRegistrationNumberOfByColor(color string) (map[string]bool, bool)
}

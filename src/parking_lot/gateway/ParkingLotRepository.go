package gateway

import (
	"container/heap"
	"errors"
	"fmt"
	"parking_lot/domain"
	"parking_lot/helpers"
	"strconv"
)

//ParkingLotRepository parking lot struct with required fields
type ParkingLotRepository struct {
	// heap of empty slots
	emptySlots helpers.IntHeap
	// max size of the parking lot
	maxParkingLotSize int
	// check if parking lot has been initialized or not
	parkingLotCreated bool
	// Map of registration number to the slot for answering queries of "slot_number_for_registration_number"
	regNoSlotMap map[string]int
	// Map of Slots to Cars for maintaining information as to which car is parked at which slot
	slotCarMap map[int]domain.Car
	// Map of Car Color to registration number Hast set used for answering queries of "slot_number_for_registration_number"
	colorRegistrationNumberMap map[string]map[string]bool
}

const (
	//WrongSizeParkingLotError error text for wrong size parking lot
	WrongSizeParkingLotError = "parking Lot of Size <= 0 cannot be created"
	//ParkingLotFullError error text for parking lot full
	ParkingLotFullError = "sorry, parking lot is full"
	//ParkingLotNotCreatedError error text for parking lot not created error
	ParkingLotNotCreatedError = "parking Lot not created"
)

//Initialize initialize parking lot
func (pL *ParkingLotRepository) Initialize(numberOfSlots int) error {
	if _, err := pL.VerifySlotInitialization(numberOfSlots); err != nil {
		return err
	}
	pL.emptySlots = helpers.IntHeap{}
	i := 1
	for i <= numberOfSlots {
		pL.emptySlots = append(pL.emptySlots, i)
		i++
	}
	heap.Init(&pL.emptySlots)
	pL.slotCarMap = map[int]domain.Car{}
	pL.colorRegistrationNumberMap = map[string]map[string]bool{}
	pL.regNoSlotMap = map[string]int{}
	pL.maxParkingLotSize = numberOfSlots
	pL.parkingLotCreated = true

	fmt.Println("Created a parking lot with " + strconv.Itoa(numberOfSlots) + " slots")
	return nil
}

//VerifySlotInitialization Verify if NumberOfSlots is correct number or not
func (pL *ParkingLotRepository) VerifySlotInitialization(numberOfSlots int) (bool, error) {
	if numberOfSlots <= 0 {
		err := errors.New(WrongSizeParkingLotError)
		return false, err
	}
	return true, nil
}

//IsParkingLotFull Verify if parking lot is full
func (pL *ParkingLotRepository) IsParkingLotFull() (bool, error) {
	if pL.emptySlots.Len() == 0 {
		err := errors.New(ParkingLotFullError)
		return true, err
	}
	return false, nil
}

//IsParkingLotCreated Verify if parking lot is created
func (pL *ParkingLotRepository) IsParkingLotCreated() (bool, error) {
	if !pL.parkingLotCreated {
		err := errors.New(ParkingLotNotCreatedError)
		return false, err
	}
	return true, nil
}

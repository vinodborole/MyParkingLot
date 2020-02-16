package mock

//ParkingLotRepository mock functions for parking lot
type ParkingLotRepository struct {
	MockInitialize               func(numberOfSlots int) error
	MockVerifySlotInitialization func(numberOfSlots int) (bool, error)
	MockIsParkingLotFull         func() (bool, error)
	MockIsParkingLotCreated      func() (bool, error)
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

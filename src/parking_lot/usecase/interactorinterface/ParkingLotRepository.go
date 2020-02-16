package interactorinterface

//ParkingLotRepository parking lot repository functions
type ParkingLotRepository interface {
	Initialize(numberOfSlots int) error
	VerifySlotInitialization(numberOfSlots int) (bool, error)
	IsParkingLotFull() (bool, error)
	IsParkingLotCreated() (bool, error)
}

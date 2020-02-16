package infra

import (
	"parking_lot/gateway"
	"parking_lot/usecase"
	"sync"
)

//ParkingLotInteract Singleton Implementation of ParkingLotRepository. Hence we use this variable to store the Instance of Parking Lot
var ParkingLotInteract *usecase.ParkingLotInteract
var once sync.Once

//GetUseCaseInteract interactor interface for parking lot functions
func GetUseCaseInteract() *usecase.ParkingLotInteract {
	once.Do(func() {
		ParkingLotRepository := gateway.ParkingLotRepository{}
		ParkingLotInteract = &usecase.ParkingLotInteract{
			ParkingLot: &ParkingLotRepository,
		}
	})
	return ParkingLotInteract
}

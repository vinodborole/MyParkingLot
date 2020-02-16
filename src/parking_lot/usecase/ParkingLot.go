package usecase

import (
	"errors"
	"fmt"
	"parking_lot/domain"
	Interact "parking_lot/usecase/interactorinterface"
	"strconv"
	"strings"
)

const (
	//NotFoundError not found error
	NotFoundError = "not found"
)

//ParkingLotInteract parking lot interact
type ParkingLotInteract struct {
	ParkingLot Interact.ParkingLotRepository
}

//Park park car in parking lot
func (sh *ParkingLotInteract) Park(car domain.Car) error {
	if _, err := sh.ParkingLot.IsParkingLotCreated(); err != nil {
		return err
	}
	// Validate if the parking lot is already full or not
	if _, err := sh.ParkingLot.IsParkingLotFull(); err != nil {
		return err
	}
	var slot int
	// Validate if the car is already parked somewhere to check mistyped input
	if slot, _ = sh.GetSlotNoForRegNo(car.GetRegNo(), false); slot == 0 {
		emptySlot := sh.ParkingLot.PopEmptySlot()
		sh.ParkingLot.MapRegNoToSlot(car.GetRegNo(), emptySlot.(int))
		sh.ParkingLot.MapSlotToCar(emptySlot.(int), car)
		sh.ParkingLot.MapColorToRegNo(car.GetColor(), car.GetRegNo())
		fmt.Println("Allocated slot number: " + strconv.Itoa(emptySlot.(int)))
		return nil
	}
	err := errors.New("Car with parkingLot registration number already parked at slot: " + strconv.Itoa(slot))
	return err
}

//Leave leave car from parking lot
func (sh *ParkingLotInteract) Leave(slot int) error {
	if _, err := sh.ParkingLot.IsParkingLotCreated(); err != nil {
		return err
	}
	// Validate if the slot has some car parked there or not
	if car, exists := sh.ParkingLot.GetCarAtSlot(slot); exists {
		sh.ParkingLot.PushEmptySlots(slot)
		sh.ParkingLot.UnmapRegNo(car.GetRegNo())
		sh.ParkingLot.UnmapRegNoFromColor(car.GetColor(), car.GetRegNo())
		sh.ParkingLot.UnmapSlot(slot)
		fmt.Println("Slot number " + strconv.Itoa(slot) + " is free")
		return nil
	}
	err := errors.New(NotFoundError)
	return err

}

//Status check status of parking lot
func (sh *ParkingLotInteract) Status() error {
	if _, err := sh.ParkingLot.IsParkingLotCreated(); err != nil {
		return err
	}
	fmt.Println("Slot No.\tRegistration No.\tColour")
	i := 1
	for i <= sh.ParkingLot.GetMaxParkingLotSize() {
		if car, exists := sh.ParkingLot.GetCarAtSlot(i); exists {
			fmt.Println(strconv.Itoa(i) + "\t" + strings.ToUpper(car.GetRegNo()) + "\t" + toCamelCase(car.GetColor()))
		}
		i++
	}
	return nil
}

//GetRegNosForCarsWithColor get registration number for cars with given color
func (sh *ParkingLotInteract) GetRegNosForCarsWithColor(color string, print bool) ([]string, error) {
	if _, err := sh.ParkingLot.IsParkingLotCreated(); err != nil {
		return []string{}, err
	}

	var regNoSlice []string
	if regNoMap, exists := sh.ParkingLot.GetRegistrationNumberOfByColor(color); exists {
		for regNo := range regNoMap {
			regNoSlice = append(regNoSlice, strings.ToUpper(regNo))
		}
	}

	if len(regNoSlice) > 0 {
		if print {
			fmt.Println(strings.Join(regNoSlice, ","))
		}
		return regNoSlice, nil
	}
	err := errors.New(NotFoundError)
	if print {
		fmt.Println(err.Error())
	}
	return regNoSlice, err
}

//GetSlotNosForCarsWithColor get slot nos for cars with given color
func (sh *ParkingLotInteract) GetSlotNosForCarsWithColor(color string) ([]int, error) {
	if _, err := sh.ParkingLot.IsParkingLotCreated(); err != nil {
		return []int{}, err
	}

	regNos, _ := sh.GetRegNosForCarsWithColor(color, false)
	var slotsString []string
	var slots []int
	for _, regNo := range regNos {
		if slot, exists := sh.ParkingLot.GetSlotForRegistrationNumber(regNo); exists {
			slotsString = append(slotsString, strconv.Itoa(slot))
			slots = append(slots, slot)
		}
	}

	if len(slots) > 0 {
		fmt.Println(strings.Join(slotsString, ","))
		return slots, nil
	}
	err := errors.New(NotFoundError)
	return slots, err

}

//GetSlotNoForRegNo get slot nos for registration number
func (sh *ParkingLotInteract) GetSlotNoForRegNo(regNo string, print bool) (int, error) {
	if _, err := sh.ParkingLot.IsParkingLotCreated(); err != nil {
		return 0, err
	}
	// Validate if there is any car with given reg no
	if slot, exists := sh.ParkingLot.GetSlotForRegistrationNumber(regNo); exists {
		if print {
			fmt.Println(slot)
		}
		return slot, nil
	}
	err := errors.New(NotFoundError)
	return 0, err
}

func toCamelCase(s string) string {
	s = strings.Trim(s, " ")
	n := ""
	capNext := true
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

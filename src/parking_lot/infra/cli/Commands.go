package cli

import (
	"errors"
	"fmt"
	"parking_lot/domain"
	"parking_lot/infra"
	"strconv"
	"strings"
)

//map of allowed commands along with the arguments to read
var allowedCommands = map[string]int{
	"create_parking_lot": 1,
	"park":               2,
	"leave":              1,
	"status":             0,
	"registration_numbers_for_cars_with_colour": 1,
	"slot_numbers_for_cars_with_colour":         1,
	"slot_number_for_registration_number":       1,
}

//ErrWrapper Wrapper Error struct
type ErrWrapper struct {
	err error
}

// Wrapper function to output the error for a given function
func (ew *ErrWrapper) do(f func() error) error {
	ew.err = f()
	if ew.err != nil {
		fmt.Println(ew.err.Error())
		return ew.err
	}
	return nil
}

const (
	//UnsupportedCommand unsupported command
	UnsupportedCommand = "unsupported command"
	//UnsupportedCommandArguments unsupported command arguments
	UnsupportedCommandArguments = "unsupported command arguments"
)

//ProcessCommand Process the command taken in from file/stdin, separate the command and arguments for command, validate the command and then do the necessary action
func ProcessCommand(command string) error {
	commandDelimited := strings.Split(command, " ")
	lengthOfCommand := len(commandDelimited)
	var arguments []string
	if lengthOfCommand < 1 {
		err := errors.New(UnsupportedCommand)
		fmt.Println(err.Error())
		return err
	} else if lengthOfCommand == 1 {
		command = commandDelimited[0]
	} else {
		command = commandDelimited[0]
		arguments = commandDelimited[1:]
	}
	// check if command is one of the allowed commands
	if numberOfArguments, exists := allowedCommands[command]; exists {

		if len(arguments) != numberOfArguments {
			err := errors.New(UnsupportedCommandArguments)
			fmt.Println(err.Error())
			return err
		}
		w := &ErrWrapper{}
		// after validation of number of arguments per command, perform the necessary command
		switch command {
		case "create_parking_lot":
			var numberOfSlots int
			var err error
			if numberOfSlots, err = strconv.Atoi(arguments[0]); err != nil {
				fmt.Println(err.Error())
				return err
			}
			return infra.GetUseCaseInteract().ParkingLot.Initialize(numberOfSlots)
		case "park":
			regNo := arguments[0]
			color := arguments[1]
			car := domain.Create(regNo, color)
			return w.do(func() error {
				return infra.GetUseCaseInteract().Park(car)
			})

		case "leave":
			var slot int
			var err error
			if slot, err = strconv.Atoi(arguments[0]); err != nil {
				fmt.Println(err.Error())
				return err
			}
			return w.do(func() error {
				return infra.GetUseCaseInteract().Leave(slot)
			})
		case "status":
			return w.do(func() error {
				return infra.GetUseCaseInteract().Status()
			})

		case "registration_numbers_for_cars_with_colour":
			color := arguments[0]
			return w.do(func() error {
				_, err := infra.GetUseCaseInteract().GetRegNosForCarsWithColor(color, true)
				return err
			})

		case "slot_numbers_for_cars_with_colour":
			color := arguments[0]
			return w.do(func() error {
				_, err := infra.GetUseCaseInteract().GetSlotNosForCarsWithColor(color)
				return err
			})

		case "slot_number_for_registration_number":
			regNo := arguments[0]
			return w.do(func() error {
				_, err := infra.GetUseCaseInteract().GetSlotNoForRegNo(regNo, true)
				return err
			})
		}
		return errors.New("not reachable code")
	}
	err := errors.New(UnsupportedCommand)
	fmt.Println(err.Error())

	return err
}

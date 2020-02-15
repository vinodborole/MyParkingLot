package domain

//Car representing car info
type Car struct {
	color              string
	registrationNumber string
}

//GetRegNo get registration number of the car
func (c Car) GetRegNo() string {
	return c.registrationNumber
}

//GetColor get color of the color
func (c Car) GetColor() string {
	return c.color
}

//Create create car object
func Create(registrationNumber string, color string) Car {
	return Car{registrationNumber: registrationNumber, color: color}
}

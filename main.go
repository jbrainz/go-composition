package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of passengers:", v.NumberOfPassengers)
	fmt.Println("Number of wheels:", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("is Electric:", c.IsElectric)
	fmt.Println("is hybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}
func main() {
	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 6,
	}

	volvoX90 := Car{
		Make:       "Volvo",
		Model:      "XC90 T8",
		Year:       2021,
		IsElectric: false,
		IsHybrid:   true,
		Vehicle:    suv,
	}
	volvoX90.show()
	fmt.Println()

	teslaModelX := Car{
		Make:       "Telsa",
		Model:      "Model X",
		Year:       2021,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}
	teslaModelX.Vehicle.NumberOfWheels = 7
	teslaModelX.show()
}

package main

import (
	"fmt"
	"math"
)

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		batteryDrain: batteryDrain,
		speed:        speed,
		battery:      100,
		distance:     0,
	}
	return car
}

func (car Car) DisplayDistance() string {
	return "Driven " + fmt.Sprint(car.distance) + " meters"
}

func (car Car) DisplayBattery() string {
	return "Battery at " + fmt.Sprint(car.battery) + "%"
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func (car Car) Drive() Car {
	if car.battery < car.batteryDrain {
		return car
	} else {
		return Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      car.battery - car.batteryDrain,
			distance:     car.distance + car.speed,
		}
	}
}

// CanFinish checks if a car is able to finish a certain track.
func (car Car) CanFinish(distance int) bool {
	td := float64(distance)
	carTimes := int(math.Ceil(td / float64(car.speed)))
	carDrain := car.batteryDrain * carTimes
	return carDrain <= car.battery
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	car2 := car.Drive()
	fmt.Println(car2.speed)
	fmt.Println(car2.batteryDrain)
	fmt.Println(car2.battery)
	fmt.Println(car2.distance)

	// fmt.Println(car.DisplayDistance())
	// trackDistance := 100
	// fmt.Println(car.CanFinish(trackDistance))

}

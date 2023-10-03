package main

import "fmt"

// Car interface represents the base car.
type Car interface {
	Drive() string
	Cost() float64
}

// BaseCar is a concrete component representing a basic car.
type BaseCar struct{}

func (c *BaseCar) Drive() string {
	return "Driving a basic car"
}

func (c *BaseCar) Cost() float64 {
	return 20000.0
}

// Decorator is the base decorator type that embeds the Car interface.
type Decorator struct {
	car Car
}

func (d *Decorator) Drive() string {
	return d.car.Drive()
}

func (d *Decorator) Cost() float64 {
	return d.car.Cost()
}

// GPSDecorator is a concrete decorator that adds GPS functionality to the car.
type GPSDecorator struct {
	Decorator
}

func NewGPSDecorator(car Car) *GPSDecorator {
	return &GPSDecorator{Decorator{car}}
}

func (d *GPSDecorator) Drive() string {
	return d.car.Drive() + " with GPS"
}

func (d *GPSDecorator) Cost() float64 {
	return d.car.Cost() + 500.0
}

// LeatherSeatsDecorator is a concrete decorator that adds leather seats to the car.
type LeatherSeatsDecorator struct {
	Decorator
}

func NewLeatherSeatsDecorator(car Car) *LeatherSeatsDecorator {
	return &LeatherSeatsDecorator{Decorator{car}}
}

func (d *LeatherSeatsDecorator) Drive() string {
	return d.car.Drive() + " with leather seats"
}

func (d *LeatherSeatsDecorator) Cost() float64 {
	return d.car.Cost() + 1000.0
}

// AlloyWheelsDecorator is a concrete decorator that adds alloy wheels to the car.
type AlloyWheelsDecorator struct {
	Decorator
}

func NewAlloyWheelsDecorator(car Car) *AlloyWheelsDecorator {
	return &AlloyWheelsDecorator{Decorator{car}}
}

func (d *AlloyWheelsDecorator) Drive() string {
	return d.car.Drive() + " with alloy wheels"
}

func (d *AlloyWheelsDecorator) Cost() float64 {
	return d.car.Cost() + 800.0
}

// Client code
func ConfigureCar() {
	baseCar := &BaseCar{}
	fmt.Printf("Base Car: %s, Cost: $%.2f\n", baseCar.Drive(), baseCar.Cost())

	gpsCar := NewGPSDecorator(baseCar)
	fmt.Printf("GPS Car: %s, Cost: $%.2f\n", gpsCar.Drive(), gpsCar.Cost())

	leatherSeatsCar := NewLeatherSeatsDecorator(gpsCar)
	fmt.Printf("Leather Seats Car: %s, Cost: $%.2f\n", leatherSeatsCar.Drive(), leatherSeatsCar.Cost())

	alloyWheelsCar := NewAlloyWheelsDecorator(leatherSeatsCar)
	fmt.Printf("Alloy Wheels Car: %s, Cost: $%.2f\n", alloyWheelsCar.Drive(), alloyWheelsCar.Cost())
}

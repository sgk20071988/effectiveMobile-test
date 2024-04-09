package db

import (
	model "effectiveMobileTest/internal/model"
)

var cars []model.Car

// Insert allows populating database
func Insert(car model.Car) {
	cars = append(cars, car)
}

// Get returns the whole database
func Get() []model.Car {
	return cars
}

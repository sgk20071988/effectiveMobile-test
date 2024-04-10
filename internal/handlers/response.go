package handlers

import (
	model "effectiveMobileTest/internal/model"
)

// JsonError is a generic error in JSON format
//
// swagger:response jsonError
type jsonError struct {
	Message string `json:"message"`
}

// CarsResponse contains a single car information
//
// swagger:response carResponse
type carResponse struct {
	Payload *model.Car `json:"car"`
}

// CarsResponse constains cars from database information
//
// swagger:response carsResponse
type carsResponse struct {
	Payload *[]model.Car `json:"cars"`
}

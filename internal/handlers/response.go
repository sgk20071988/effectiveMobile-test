package handlers

import (
	model "effectiveMobileTest/internal/model"
)

type jsonError struct {
	Message string `json:"message"`
}

type carResponse struct {
	Payload *model.Car `json:"car"`
}

type carsResponse struct {
	Payload *[]model.Car `json:"cars"`
}

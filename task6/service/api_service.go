package service

import (
	"../logger"
	"../repository"
)

// Services
var Apartment ApartmentService

// Logger
var log = logger.Get()

func Init() {
	// Init services
	Apartment = &apartmentService{repository.Apartment}
}

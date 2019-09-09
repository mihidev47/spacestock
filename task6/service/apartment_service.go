package service

import (
	"../model"
	"../repository"
)

type ApartmentService interface {
	GetApartment() interface{}
	AddApartment(name string, address string) interface{}
	UpdateApartment(Id string, name string, address string) interface{}
	DeleteApartment(Id string) interface{}
}

type apartmentService struct {
	aptRepo   repository.AptRepository
}

func (s *apartmentService) GetApartment() interface{} {
	// Get timestamp
	aptRepo := repository.Apartment
	
	// Compose response
	result, _ := aptRepo.GetAll()
	// result["result_jwt"] = response.NewToken(token, expiredAt)

	return &result
}

func (s *apartmentService) AddApartment(name string, address string) interface{} {
	aptRepo := repository.Apartment
	result := make(map[string]interface{})
	// Get timestamp
	newApt := model.NewApartment(name, address)
	
	_, err := aptRepo.Insert(&newApt)
	if err != nil {
		result["status"] = "Failed"
		result["message"] = err
		return result
	}
	// Compose response
	result["status"] = "Success"
	// result["result_jwt"] = response.NewToken(token, expiredAt)

	return &result
}


func (s *apartmentService) UpdateApartment(Id string, name string, address string) interface{} {
	aptRepo := repository.Apartment
	result := make(map[string]interface{})
	// Get timestamp
	newApt := model.NewApartment(name, address)
	
	err := aptRepo.Update(Id, &newApt)
	if err != nil {
		result["status"] = "Failed"
		result["message"] = err
		return result
	}
	// Compose response
	result["status"] = "Success"
	// result["result_jwt"] = response.NewToken(token, expiredAt)

	return &result
}

func (s *apartmentService) DeleteApartment(Id string) interface{} {
	aptRepo := repository.Apartment
	result := make(map[string]interface{})
	// Get timestamp
	_ = aptRepo.Delete(Id)
	// Compose response
	result["status"] = true
	// result["result_jwt"] = response.NewToken(token, expiredAt)

	return &result
}

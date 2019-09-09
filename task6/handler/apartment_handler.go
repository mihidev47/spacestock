package handler

import (
	"net/http"
	"../request"
	"../response"
	"../service"
	"../util"
	"github.com/gorilla/mux"
)

func GetApartment(r *http.Request) (*response.Success, error) {
	// @todo Decrypt password
	// @todo Validate password
	// Register from foody service
	// Call service
	resBody := service.Apartment.GetApartment()

	// Compose response
	return response.NewSuccess(resBody), nil
}

func AddApartment(r *http.Request) (*response.Success, error) {
	// Parse request body
	var req request.Apartment
	err := parseJSON(r, &req)
	if err != nil {
		return nil, util.NewError("400")
	}
	// @todo Decrypt password
	// @todo Validate password
	// Call service
	resBody := service.Apartment.AddApartment(req.Name, req.Address)

	// Compose response
	return response.NewSuccess(resBody), nil
}

func UpdateApartment(r *http.Request) (*response.Success, error) {
	// Parse request body
	var req request.Apartment

	err := parseJSON(r, &req)
	if err != nil {
		return nil, util.NewError("400")
	}

	Id := mux.Vars(r)["id"]

	// @todo Decrypt password
	// @todo Validate password
	// Register from foody service
	// Call service
	resBody := service.Apartment.UpdateApartment(Id, req.Name, req.Address)

	// Compose response
	return response.NewSuccess(resBody), nil
}

func DeleteApartment(r *http.Request) (*response.Success, error) {
	// Parse request body
	Id := mux.Vars(r)["id"]
	// @todo Decrypt password
	// @todo Validate password
	// Register from foody service
	// Call service
	resBody := service.Apartment.DeleteApartment(Id)

	// Compose response
	return response.NewSuccess(resBody), nil
}

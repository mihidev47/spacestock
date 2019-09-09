package router

import (
	"../flags"
	"../handler"
)

// routeAPI configure request routing in API. Handlers must be defined in handler package
func routeAPI(r Router) {
	r.HandleREST("/apartment", handler.GetApartment, flags.ACLAuthenticatedAnonymous).Methods("GET")
	r.HandleREST("/apartment", handler.AddApartment, flags.ACLAuthenticatedAnonymous).Methods("POST")
	r.HandleREST("/apartment/{id}", handler.UpdateApartment, flags.ACLAuthenticatedAnonymous).Methods("PUT")
	r.HandleREST("/apartment/{id}", handler.DeleteApartment, flags.ACLAuthenticatedAnonymous).Methods("DELETE")
}

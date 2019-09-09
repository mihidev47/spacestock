package model

type Apartment struct {
	Id        	string         `db:"id"`
	Name   		string         `db:"name"`
	Address     string         `db:"address"`
}

// NewFoody construct Foody instance
func NewApartment(name string, address string) Apartment {
	return Apartment{
		Name:   name,
		Address:  address,
	}
}

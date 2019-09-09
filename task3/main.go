package main

import (
	"errors"
	"fmt"
)

// Person Struct
type Person struct {
	Nama   string
	Jk string
	Umur    int
}

func main() {
	jon := NewPerson().Name("Jon Snow").Gender("Male").Age(20)
	fmt.Println(jon)
}

// NewPerson Return New
func NewPerson() *Person {
	var p Person
	return &p
}

// Name Func
func (p *Person) Name(names string) *Person {
	p.Nama = names
	return p
}

// Gender Func
func (p *Person) Gender(gender string) *Person {
	p.Jk = gender
	return p
}

// Age Func
func (p *Person) Age(age int) *Person {
	p.Umur = age
	return p
}

// Validate Func
func (p Person) Validate() error {
	if p.Nama == "" {
		return errors.New("Name cannot be empty")
	}

	if p.Jk != "Male" && p.Jk != "Female" {
		return errors.New("Gender is either Male or Female")
	}

	if p.Umur < 0 {
		return errors.New("There is no such thing as negative age")
	}

	return nil
}

package main

import (
	"testing"
)

func TestValidate(t *testing.T) {
	var p Person
	p.Name("Jon Snow").Gender("Male").Age(20)
	err := p.Validate()
	if err != nil {
		t.Errorf(err.Error())
	}
}

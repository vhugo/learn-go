package main

import "fmt"

type Car interface {
	NumDoors() int
}

const (
	_ = iota
	LuxuryCarType
	FamilyCarType
)

type CarFactory struct{}

func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil

	case FamilyCarType:
		return new(FamilyCar), nil

	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized\n", v)
	}
}

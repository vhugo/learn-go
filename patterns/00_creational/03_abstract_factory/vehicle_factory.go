package main

import "fmt"

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

const (
	_ = iota
	CarFactoryType
	MotorbikeFactoryType
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil

	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil

	default:
		return nil, fmt.Errorf("Factory with id %d not recognized\n", f)
	}
}

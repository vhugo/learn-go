package main

import "fmt"

type Motorbike interface {
	GetMotorbikeType() int
}

const (
	_ = iota
	SportMotorbikeType
	CruiseMotorbikeType
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil

	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil

	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized\n", v)
	}
}

// Prototype
//
// Aims to have an object or a set of objects that is already created at
// compilation time, but which you can clone as many times as you want at
// runtime. The key difference between this and a Builder pattern is that
// objects are cloned for the user instead of building them at runtime. You can
// also build a cache-like solution, storing information using a prototype.
//
// The main objective is to avoid repetitive object creation.
//
// - Maintain a set of objects that will be cloned to create new instances
//
// - Provide a default value of some type to start working on top of it
//
// - Free CPU of complex object initialization to take more memory resources
//
package main

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	_ = iota
	White
	Black
	Blue
)

func GetShirtsCloner() ShirtCloner {
	return new(ShirtsCache)
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(c int) (ItemInfoGetter, error) {
	switch c {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n",
		s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

// Decorator
//
// The Decorator pattern is pretty simple, but, for instance, it provides a lot
// of benefits when working with legacy code. It allows you to decorate an
// already existing type with more functional features without actually touching
// it. The Decorator type implements the same interface of the type it
// decorates, and stores an instance of that type in its members. This way, you
// can stack as many decorators (dolls) as you want by simply storing the old
// decorator in a field of the new one.
//
// When you need to add functionality to some code that you don't have access
// to, or you don't want to modify to avoid a negative effect on the code, and
// follow the open/close principle (like legacy code)
//
// When you want the functionality of an object to be created or altered
// dynamically, and the number of features is unknown and could grow fast
//
package main

import (
	"errors"
	"fmt"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type Meat struct {
	Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field " +
			"of the Meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field " +
			"of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "onion"), nil
}

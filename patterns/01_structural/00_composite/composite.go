// Composite
//
// Favors composition (commonly defined as a has a relationship) over
// inheritance (an is a relationship). You will create hierarchies and trees of
// objects. Objects have different objects with their own fields and methods
// inside them. This approach is very powerful and solves many problems of
// inheritance and multiple inheritances.
//
package main

const (
	_ = iota
	Training
	Swimming
	Eating
)

type Athlete struct{}

func (a *Athlete) Train() int {
	return Training
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func() int
}

func Swim() int {
	return Swimming
}

type Animal struct{}

func (r *Animal) Eat() int {
	return Eating
}

type Shark struct {
	Animal
	Swim func() int
}

type Swimmer interface {
	Swim() int
}

type Trainer interface {
	Train() int
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() int {
	return Swimming
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

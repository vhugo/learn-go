// Requirements and acceptance criteria
//
// We'll have an athlete and a swimmer. We will also have an animal and a fish.
// The Swimmer and the Fish methods must share the code. The athlete must train,
// and the animal must eat:
//
// - We must have an Athlete struct with a Train method
//
// - We must have a Swimmer with a Swim method
//
// - We must have an Animal struct with an Eat method
//
// - We must have a Fish struct with a Swim method that is shared with the
// Swimmer, and not have inheritance or hierarchy issues
//
package main

import "testing"

func TestSwimmer(t *testing.T) {
	swimmerA := CompositeSwimmerA{MySwim: Swim}
	if swimmerA.MySwim() != Swimming {
		t.Error("Swimmer should swim")
	}

	if swimmerA.MyAthlete.Train() != Training {
		t.Error("Swimmer should train")
	}

	swimmerB := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}
	if swimmerB.Swim() != Swimming {
		t.Error("Swimmer should swim")
	}

	if swimmerB.Train() != Training {
		t.Error("Swimmer should train")
	}
}

func TestFish(t *testing.T) {
	fish := Shark{Swim: Swim}
	if fish.Swim() != Swimming {
		t.Error("Fish should swim")
	}

	if fish.Eat() != Eating {
		t.Error("Fish should eat")
	}
}

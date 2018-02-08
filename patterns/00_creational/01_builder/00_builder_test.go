// Requirements and acceptance criteria
//
// As far as we have described, we must dispose of a Builder of type Car and
// Motorbike and a unique director called ManufacturingDirector to take builders
// and construct products. So the requirements for a Vehicle builder example
// would be the following:
//
// I must have a manufacturing type that constructs everything that a vehicle
// needs When using a car builder, the VehicleProduct with four wheels, five
// seats, and a structure defined as Car must be returned
//
// When using a motorbike builder, the VehicleProduct with two wheels, two
// seats, and a structure defined as Motorbike must be returned
//
// A VehicleProduct built by any BuildProcess builder must be open to
// modifications
//
package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n",
			motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n",
			motorbike.Structure)
	}

}

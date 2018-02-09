// Builder
//
// The Builder design pattern has been commonly described as the relationship
// between a director, a few Builders, and the product they build.
//
// A Builder design pattern tries to:
//
// - Abstract complex creations so that object creation is separated from the
// object user
//
// - Create an object step by step by filling its fields and creating the
// embedded objects
//
// - Reuse the object creation algorithm between many objects
//
package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.Build()
}

type BikeBuilder struct {
	v VehicleProduct
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 1
	return c
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Motorbike"
	return c
}

func (c *BikeBuilder) Build() VehicleProduct {
	return c.v
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.Build()
}

package truck

import "ejercicio-1/vehicle"

// Type assertion TruckBuilder must implement Builder interface
var _ vehicle.Builder = &TruckBuilder{}

var _ Builder = &TruckBuilder{}

// Builder defines the behavior for truck builders
type Builder interface {
	// Take behavior from vehicle builder
	vehicle.Builder

	// Define specific behavior for trucks builder
	SetMaxLoad(maxLoad int)
	BuildTruck() *Truck
}

// TruckBuilder is an specific implementation to build truck vehicles
type TruckBuilder struct {
	*Truck
}

// NewTruckBuilder creates a new instance of TruckBuilder
func NewTruckBuilder() TruckBuilder {
	return TruckBuilder{Truck: &Truck{}}
}

// Reset returns the truck to an empty state
func (t *TruckBuilder) Reset() {
	t.BasicVehicle.Reset()
	t.maxLoad = 0
}

// SetMaxLoad sets max load for the truck
func (t *TruckBuilder) SetMaxLoad(maxLoad int) {
	t.maxLoad = maxLoad
}

// BuildTruck define builds the truck object
func (t *TruckBuilder) BuildTruck() *Truck {
	// Set default if required info were not provided
	t.BasicVehicle.ApplyDefaults()

	// return built object
	return t.Truck
}

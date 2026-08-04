package truck

import "ejercicio-1/vehicle"

// Type assertion Truck must implement Vehicle interface
var _ vehicle.Vehicle = &Truck{}

// Type assertion Truck must implement TruckVehicle interface
var _ TruckVehicle = &Truck{}

// Truck is a truck vehicle
type Truck struct {
	// extend basic vehicle behavior to truck
	vehicle.BasicVehicle

	// Specific properties for truck
	maxLoad int
}

// GetMaxLoad returns the maximum load the truck handle
func (t *Truck) GetMaxLoad() int {
	return t.maxLoad
}

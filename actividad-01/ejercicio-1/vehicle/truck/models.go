package truck

import "ejercicio-1/vehicle"

// TruckVehicle defines the specific behavior for trucks
type TruckVehicle interface {
	// Take vehicle definition to Truck
	vehicle.Vehicle

	// GetMaxLoad returns the maximum load the truck handle
	GetMaxLoad() int
}

package sedan

import "ejercicio-1/vehicle"

// Type interface Sedan must implement Vehicle
var _ SedanVehicle = &Sedan{}

// Sedan is a vehicle
type Sedan struct {
	// Sedan is a basic vehicle do not need to re-implement Vehicle interface
	vehicle.BasicVehicle
}

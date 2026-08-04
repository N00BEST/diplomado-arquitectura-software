package main

import (
	"fmt"

	"ejercicio-1/vehicle"
	"ejercicio-1/vehicle/sedan"
	"ejercicio-1/vehicle/truck"
)

func main() {
	// Initialize builders
	truckBuilder := truck.NewTruckBuilder()
	sedanBuilder := sedan.NewSedanBuilder()

	// Create truck1 as default truck
	truck1 := truckBuilder.BuildTruck()
	fmt.Printf("truck1: %+v\n", truck1)

	// Create truck2 I need to reset to create a new one from scrach
	truckBuilder.Reset()
	truckBuilder.SetColor(vehicle.RedColor)
	truckBuilder.SetMaxLoad(2000)
	truck2 := truckBuilder.BuildTruck()
	fmt.Printf("truck2: %+v\n", truck2)

	// Create bmwSedan as default sedan
	bmwSedan := sedanBuilder.BuildSedan()
	fmt.Printf("bmwSedan: %+v\n", bmwSedan)

	// Create audiSedan I need to reset to create a new one from scrach
	sedanBuilder.Reset()
	sedanBuilder.SetColor(vehicle.WhiteColor)
	sedanBuilder.SetInterior(vehicle.LeatherInsiderType)
	audiSedan := sedanBuilder.BuildSedan()
	fmt.Printf("audiSedan: %+v\n", audiSedan)
}

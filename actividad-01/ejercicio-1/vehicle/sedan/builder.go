package sedan

import "ejercicio-1/vehicle"

// Type assertion SedanBuilder must implement Builder interface
var _ vehicle.Builder = &SedanBuilder{}

var _ Builder = &SedanBuilder{}

// Builder defines the expected behavior for Sedan builder
type Builder interface {
	// Take builder definition from vehicle definition
	vehicle.Builder

	// Define specific behavior for sedan builder
	BuildSedan() *Sedan
}

// SedanBuilder is an specific implementation to build sedan vehicles
type SedanBuilder struct {
	*Sedan
}

// NewSedanBuilder creates a new instance of SedanBuilder
func NewSedanBuilder() SedanBuilder {
	return SedanBuilder{Sedan: &Sedan{}}
}

// BuildSedan returns the sedan built with the provided information
func (s *SedanBuilder) BuildSedan() *Sedan {
	// Set default if required info for the Vehicle were not provided
	s.BasicVehicle.Build()

	// return built object
	return s.Sedan
}

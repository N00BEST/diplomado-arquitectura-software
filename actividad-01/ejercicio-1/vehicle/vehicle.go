package vehicle

// Type assertion BasicVehicle implements Vehicle
var _ Vehicle = &BasicVehicle{}

// BasicVehicle contains the basic elements for a vehicle definition
type BasicVehicle struct {
	motorType   MotorType
	color       Color
	wheels      int
	soundSystem SoundSystem
	interior    InteriorType
	solarPanel  bool
	gps         bool
}

// GetMotorType returns the motor type of the vehicle
func (b *BasicVehicle) GetMotorType() MotorType {
	return b.motorType
}

// GetColor returns the color of the vehicle
func (b *BasicVehicle) GetColor() Color {
	return b.color
}

// GetWheels returns the amount of wheels of the vehicle
func (b *BasicVehicle) GetWheels() int {
	return b.wheels
}

// GetSoundSystem returns the sound system of the vehicle
func (b *BasicVehicle) GetSoundSystem() SoundSystem {
	return b.soundSystem
}

// GetInterior returns the interior type of the vehicle
func (b *BasicVehicle) GetInterior() InteriorType {
	return b.interior
}

// GetGps returns whether the vehicle has gps
func (b *BasicVehicle) GetGps() bool {
	return b.gps
}

// GetSolarPanel returns whether the vehicle has solar panel
func (b *BasicVehicle) GetSolarPanel() bool {
	return b.solarPanel
}

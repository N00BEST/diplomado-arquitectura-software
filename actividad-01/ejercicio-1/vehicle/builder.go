package vehicle

// Type assertion BasicVehicle must implement Builder interface
var _ Builder = &BasicVehicle{}

// Reset returns an empty Vehicle
func (b *BasicVehicle) Reset() {
	*b = BasicVehicle{}
}

// SetMotor sets the motor type to the vehicle
func (b *BasicVehicle) SetMotor(motorType MotorType) {
	b.motorType = motorType
}

// SetColor sets the color to the vehicle
func (b *BasicVehicle) SetColor(color Color) {
	b.color = color
}

// SetWheels sets the number of wheels to the vehicle
func (b *BasicVehicle) SetWheels(wheels int) {
	b.wheels = wheels
}

// SetSoundSystem sets the sound system to the vehicle
func (b *BasicVehicle) SetSoundSystem(soundType SoundSystem) {
	b.soundSystem = soundType
}

// SetInterior sets the interior to the vehicle
func (b *BasicVehicle) SetInterior(interiorType InteriorType) {
	b.interior = interiorType
}

// SetSolarPanel sets the solar panel to the vehicle
func (b *BasicVehicle) SetSolarPanel(solarPanel bool) {
	b.solarPanel = solarPanel
}

// SetGps sets the gps to the vehicle
func (b *BasicVehicle) SetGps(gps bool) {
	b.gps = gps
}

// Build returns the object created with all parameters
func (b *BasicVehicle) Build() Vehicle {
	// Set default if required info were not provided
	b.setDefaults()

	// return built object
	return b
}

// setDefaults assign the default values of the Vehicle
func (s *BasicVehicle) setDefaults() {
	if s.color == "" {
		s.color = GrayColor
	}

	if s.wheels == 0 {
		s.wheels = 4
	}

	if s.soundSystem == "" {
		s.soundSystem = MonoSoundSystem
	}

	if s.interior == "" {
		s.interior = FabricInsiderType
	}

}

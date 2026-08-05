package vehicle

// Builder defines the interface with the expected behavior
// to build a Sedan vehicle
type Builder interface {
	// Reset returns an empty Vehicle
	Reset()
	// SetMotor sets the motor type to the vehicle
	SetMotor(motorType MotorType)
	// SetColor sets the color to the vehicle
	SetColor(color Color)
	// SetWheels sets the number of wheels to the vehicle
	SetWheels(wheels int)
	// SetSoundSystem sets the sound system to the vehicle
	SetSoundSystem(soundType SoundSystem)
	// SetInterior sets the interior to the vehicle
	SetInterior(insiderType InteriorType)
	// SetSolarPanel sets the solar panel to the vehicle
	SetSolarPanel(solarPanel bool)
	// SetGps sets the gps to the vehicle
	SetGps(gps bool)
}

// Vehicle defines the expected behavior for vehicles
type Vehicle interface {
	// GetMotorType returns the motor type of the vehicle
	GetMotorType() MotorType
	// GetColor returns the color of the vehicle
	GetColor() Color
	// GetWheels returns the amount of wheels of the vehicle
	GetWheels() int
	// GetSoundSystem returns the sound system of the vehicle
	GetSoundSystem() SoundSystem
	// GetInterior returns the interior type of the vehicle
	GetInterior() InteriorType
	// GetGps returns whether the vehicle has gps
	GetGps() bool
	// GetSolarPanel returns whether the vehicle has solar panel
	GetSolarPanel() bool
}

// MotorType is a custom type to define the supported motor types
type MotorType string

// Supported motor types
const (
	ManualMotorType    MotorType = "manual"
	AutomaticMotorType MotorType = "automatic"
)

// SoundSystem is a custom type to define the supported sound types
type SoundSystem string

// Supported sound systems
const (
	SurroundSoundSystem SoundSystem = "surround"
	MonoSoundSystem     SoundSystem = "mono"
)

// InteriorType is a custom type to define the supported vehicle interior
type InteriorType string

// Supported insider type
const (
	FabricInsiderType  InteriorType = "fabric"
	LeatherInsiderType InteriorType = "leather"
)

// Color is a custom type to define the supported colors
type Color string

// Supported colors
const (
	WhiteColor Color = "white"
	BlackColor Color = "black"
	BlueColor  Color = "blue"
	RedColor   Color = "red"
	GrayColor  Color = "gray"
)

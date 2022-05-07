package app

type LocationProvider interface {
	GetLocation() string
}

type DeviceManager interface {
	Init() error
	SetLocation(str string) error
	ResetLocation() error
}

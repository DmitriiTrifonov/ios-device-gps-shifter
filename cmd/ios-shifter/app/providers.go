package app

type LocationProvider interface {
	GetLocation() string
}

type DeviceManager interface {
	GetDevice() string
	LoadImages(uuid string) error
	SetLocation(str string) error
	ResetLocation() error
}

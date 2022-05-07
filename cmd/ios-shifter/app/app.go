package app

import "log"

type Application struct {
	loc           LocationProvider
	deviceManager DeviceManager
}

func NewApplication(loc LocationProvider, manager DeviceManager) *Application {
	return &Application{
		loc:           loc,
		deviceManager: manager,
	}
}

func (a *Application) Start() {
	err := a.deviceManager.Init()
	if err != nil {
		// TODO: Add retries and proper error handling
		log.Fatal(err)
	}
	a.loop()
}

func (a *Application) loop() {
	var locStr string
	for {
		locStr = a.loc.GetLocation()
		err := a.deviceManager.SetLocation(locStr)
		if err != nil {
			// TODO: Add error handling
			log.Fatal(err)
		}
	}
}

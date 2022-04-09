package location

import gpssh "github.com/DmitriiTrifonov/gps-shifter"

type Location struct {
	port            string
	initialLocation string
	currentLocation string
	receiver        *gpssh.Receiver
	shifter         *gpssh.Shifter
	shiftChan       chan gpssh.Vector2D
}

func (l *Location) GetLocation() string {
	strCoords := <-l.shiftChan
	return strCoords.String()
}

func NewLocation(port, initialLocation string) (*Location, error) {
	receiver, err := gpssh.NewReceiver(port, 512, 9600)
	if err != nil {
		return nil, err
	}

	initCoords, err := gpssh.ParseVector2DFromString(initialLocation)
	if err != nil {
		return nil, err
	}

	shifter := gpssh.NewShifter(receiver, initCoords)

	shiftChannel := make(chan gpssh.Vector2D, 1000)

	loc := &Location{
		port:            port,
		initialLocation: initialLocation,
		currentLocation: "",
		receiver:        receiver,
		shifter:         shifter,
		shiftChan:       shiftChannel,
	}

	go loc.shifter.Shift(loc.shiftChan, initCoords)

	return loc, nil
}

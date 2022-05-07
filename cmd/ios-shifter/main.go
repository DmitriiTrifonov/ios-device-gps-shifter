package main

import (
	"flag"
	"fmt"
	"ios-device-gps-shifter/cmd/ios-shifter/app"
	"ios-device-gps-shifter/internal/device"
	"ios-device-gps-shifter/internal/location"
	"log"
)

func main() {
	fmt.Println("Hello")
	portPtr := flag.String("port", "/dev/ttyS1", "sets a serial port")
	initLocationPtr := flag.String("initial", "56.946879,24.121411", "sets a serial port")
	imagesPathPtr := flag.String("images", "/home/", "sets images path")
	flag.Parse()

	loc, err := location.NewLocation(*portPtr, *initLocationPtr)
	if err != nil {
		log.Fatal(err)
	}

	man := device.NewManager(*imagesPathPtr)

	application := app.NewApplication(loc, man)

	application.Start()
}

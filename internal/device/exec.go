package device

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	deviceIDBin     = "idevice_id"
	imageMounterBin = "ideviceimagemounter"
	setLocationBin  = "idevicesetlocation"
)

func (m *Manager) getDeviceUUID() (string, error) {
	cmd := exec.Command(deviceIDBin)

	data, err := cmd.Output()
	if err != nil {
		return "", err
	}

	dataStr := strings.Split(string(data), " ")
	if len(dataStr) != 2 {
		return "", fmt.Errorf("cannot parse device id from %s", string(data))
	}

	m.uuid = dataStr[0]
	return dataStr[0], nil
}

func (m *Manager) loadImages() error {
	var dmg, signature string
	entries, err := os.ReadDir(m.imagePath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		filename := entry.Name()
		if strings.HasSuffix(filename, ".dmg") {
			dmg = m.imagePath + "/" + filename
			continue
		}
		if strings.HasSuffix(filename, ".signature") {
			signature = m.imagePath + "/" + filename
			continue
		}

	}

	log.Println("dmg path:", dmg)
	log.Println("signature path:", signature)

	listCommand := exec.Command(imageMounterBin, "-u", m.uuid, "-l")
	listCommandData, err := listCommand.Output()
	log.Println("list command output:", string(listCommandData))
	err = checkDeviceLocked(listCommandData)
	if err != nil {
		return err
	}

	if checkForNotMounted(listCommandData) {
		log.Println("images already mounted", string(listCommandData))
		return nil
	}

	// Not initialized device
	//Uploading DeveloperDiskImage.dmg
	//done.
	//Mounting...
	//Done.
	//Status: Complete

	cmd := exec.Command(imageMounterBin, "-u", m.uuid, dmg, signature)
	cmdData, err := cmd.Output()
	err = checkDeviceLocked(cmdData)
	if err != nil {
		return err
	}

	log.Println("mounting command output:", string(cmdData))

	if !strings.Contains(string(cmdData), "Status: Complete") {
		return fmt.Errorf("cannot parse state: %s", string(cmdData))
	}
	return nil
}

func checkForNotMounted(data []byte) bool {
	return strings.Contains(string(data), "ImageSignature[1]:")
}

func checkDeviceLocked(data []byte) error {
	if strings.Contains(string(data), "Error: DeviceLocked") {
		return fmt.Errorf("device is locked: unlock the device to proceed")
	}
	return nil
}

func (m *Manager) setLocation(input string) error {
	coords := strings.Split(input, " ")
	if len(coords) != 2 {
		return fmt.Errorf("cannot parse input: %s", input)
	}

	cmd := exec.Command(setLocationBin, "-u", m.uuid, "--", coords[0], coords[1])
	out, err := cmd.Output()
	if err != nil {
		log.Println("output:", out)
		return err
	}
	log.Println("set coordinates:", input)
	return nil
}

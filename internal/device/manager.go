package device

import "log"

type State int

const (
	Unknown State = iota
	Init
	Ready
	Operated
)

type Manager struct {
	state     State
	imagePath string
	uuid      string
}

func NewManager(imagesPath string) *Manager {
	return &Manager{
		state:     Unknown,
		imagePath: imagesPath,
	}
}

func (m *Manager) Init() error {
	log.Println("initializing device manager")
	m.state = Init
	log.Println("device state:", m.GetState().String())
	deviceUUID, err := m.getDeviceUUID()
	if err != nil {
		return err
	}
	log.Printf("got device: %s\n", deviceUUID)
	err = m.loadImages()
	if err != nil {
		return err
	}
	m.state = Ready
	log.Println("device state:", m.GetState().String())
	return nil
}

func (m *Manager) GetState() State {
	return m.state
}

func (m *Manager) SetLocation(locStr string) error {
	return m.setLocation(locStr)
}

func (m *Manager) ResetLocation() error {
	return nil
}

func (s State) String() string {
	switch s {
	case Init:
		return "StateInit"
	case Ready:
		return "StateReady"
	case Operated:
		return "StateOperated"
	default:
		return "StateUnknown"
	}
}

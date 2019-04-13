package entities

type Port struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type PortRepository interface {
	GetPort(id string) (Port, error)
	GetAllPorts(limit, offset uint) []Port
	CreatePort(Port) error
	UpdatePort(Port) error
	CreateOrUpdate(Port) error
}

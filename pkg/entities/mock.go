package entities

var MockPort = Port{ID: "AEAJM", Name: "Ajman", City: "Ajman", Country: "United Arab Emirates", Coordinates: []float64{55.5136433, 25.4052165}, Province: "Ajman", Timezone: "Asia/Dubai", Unlocs: []string{"AEAJM"}, Code: "52000"}

func MockPorts(len int) []Port {
	ports := make([]Port, 0, len)
	for i := 0; i < len; i++ {
		ports = append(ports, MockPort)
	}
	return ports
}

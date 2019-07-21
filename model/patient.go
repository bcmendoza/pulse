package model

type Patient struct {
	Type     string            `json:"type"`
	Name     string            `json:"name"`
	Children map[string]Metric `json:"children"`
	Stream   Stream            `json:"stream"`
}

type PatientKey struct {
	Department string
	Patient    string
}

func (h *Hospital) AddPatient(department, patient string) {
	h.Lock()
	defer h.Unlock()

	p := Patient{
		Type:     "patient",
		Name:     patient,
		Children: make(map[string]Metric),
		Stream: Stream{
			UnitType:   "%",
			History:    make([]Pulse, 0),
			Lower:      0,
			Upper:      100,
			Thresholds: []float64{33.33, 66.66},
		},
	}
	h.Children[department].Children[patient] = p
}
